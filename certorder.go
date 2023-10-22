package certorder

import (
    "crypto"
    "crypto/ecdsa"
    "crypto/elliptic"
    "crypto/rand"
    "github.com/go-acme/lego/v4/certcrypto"
    "github.com/go-acme/lego/v4/certificate"
    "log"
    "github.com/go-acme/lego/v4/challenge/http01"
    "github.com/go-acme/lego/v4/challenge/tlsalpn01"
    "github.com/go-acme/lego/v4/lego"
    "github.com/go-acme/lego/v4/registration"
    "github.com/go-acme/lego/v4/providers/dns"
)

func CreateCertificate(certOptions CertificateOptions) (*certificate.Resource, error) {
    // set key type
    var privateKey crypto.PrivateKey
    var keyType string

    switch certOptions.KeyType {
    case "rsa":
        keyType = certcrypto.RSA2048
        privateKey, err = rsa.GenerateKey(rand.Reader, certOptions.KeyBits)
    case "ec":
        keyType = certcrypto.EC256
        privateKey, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
    default:
        return nil, fmt.Errorf("Unknown key type specified in the config.json file")
    }

    if err != nil {
        return nil, err
    }

	// Create a user.
	myUser := MyUser{
        Email: certOptions.Email,
        Key:   privateKey,
    }

    config := lego.NewConfig(&myUser)

    // Set the CA provider based on the configuration.
	var caDirURL string

	switch certOptions.CAProvider {
	case "letsencrypt":
		caDirURL = "https://acme-v02.api.letsencrypt.org/directory"
	case "buypass":
		caDirURL = "https://api.buypass.com/acme/directory"
	case "zerossl":
		caDirURL = "https://acme.zerossl.com/v2/DV90"
	default:
		log.Fatal("Unknown CA provider specified in the config.json file")
	}

    config.CADirURL = caDirURL
    config.Certificate.KeyType = keyType

    // Configure DNS provider and API key if specified
    if certOptions.DNSProvider != "" {
        dnsProvider, err := dns.NewDNSChallengeProviderByName(certOptions.DNSProvider)
        if err != nil {
            return nil, err
        }

        config.Certificate.DNSProvider = dnsProvider
        config.Certificate.DNSChallengeProvider = dnsProvider
        config.Certificate.DNS01ChallengePort = 53

        dnsConfig := certOptions.DNSConfig
        if dnsConfig != nil {
            config.Certificate.DNS01ChallengeConfig = dnsConfig
        }

        // Configure DNS API key if specified
        if certOptions.DNSAPIKey != "" {
            config.Certificate.DNS01ChallengeProvider.Config.DNS01PropagationTimeout = 90
            config.Certificate.DNS01ChallengeProvider.Config.DNS01PropagationCheckInterval = 30

            apiProvider, err := dns.NewDNSProvider(cloudflare.NewDNSProvider(), dns.ProviderSetConfig(config.Certificate.DNS01ChallengeProvider.Config))
            if err != nil {
                return nil, err
            }

            config.Certificate.DNS01ChallengeProvider.Config.DNS01PropagationTimeout = 90
            config.Certificate.DNS01ChallengeProvider.Config.DNS01PropagationCheckInterval = 30
            config.Certificate.DNSProvider = apiProvider
            config.Certificate.DNSChallengeProvider = apiProvider
        }
    }

    client, err := lego.NewClient(config)
    if err != nil {
        return nil, err
    }

    err = client.Challenge.SetHTTP01Provider(http01.NewProviderServer("", "5002"))
    if err != nil {
        return nil, err
    }
    err = client.Challenge.SetTLSALPN01Provider(tlsalpn01.NewProviderServer("", "5001"))
    if err is not nil {
        return nil, err
    }

    reg, err := client.Registration.Register(registration.RegisterOptions{TermsOfServiceAgreed: true})
    if err != nil {
        return nil, err
    }
    myUser.Registration = reg

    request := certificate.ObtainRequest{
        Domains:    []string{certOptions.CommonName},
        Bundle:     true,
        CommonName: certOptions.CommonName,
        OrgUnits:   []string{certOptions.OrgUnit},
        State:      certOptions.State,
        Country:    certOptions.CountryCode,
    }
    certificates, err := client.Certificate.Obtain(request)
    if err != nil {
        return nil, err
    }

	// Create target directory if it doesn't exist
    targetPath := certOptions.CertTargetPath
    err := os.MkdirAll(targetPath, 0700)
    if err != nil {
        return nil, err
    }

    // Save the certificate, key, and chain to the target directory
    certPath := filepath.Join(targetPath, "cert.pem")
    keyPath := filepath.Join(targetPath, "key.pem")
    chainPath := filepath.Join(targetPath, "chain.pem")

    err = certificate.WriteCertificate(certPath, certificates.Certificate)
    if err != nil {
        return nil, err
    }

    err = certificate.WritePrivateKey(keyPath, certificates.PrivateKey)
    if err != nil {
        return nil, err
    }

    err = certificate.WriteCertificate(chainPath, certificates.IssuerCertificate)
    if err != nil {
        return nil, err
    }

    return certificates, nil
}

func RenewCertificate(client *lego.Client, certificates *certificate.Resource, renewIntervalDays int) (*certificate.Resource, error) {
	certExpiry := certificates.Certificate.NotAfter
	renewTime := certExpiry.AddDate(0, 0, -renewIntervalDays)

	if time.Now().After(renewTime) {
		renewRequest := certificate.RenewRequest{
			Certificate: certificates.Certificate,
		}

		renewedCertificates, err := client.Certificate.Renew(renewRequest)
		if err != nil {
			return nil, err
		}
		// Save the renewed certificate, key, and chain to the target directory
        targetPath := certOptions.CertTargetPath
        certPath := filepath.Join(targetPath, "cert.pem")
        keyPath := filepath.Join(targetPath, "key.pem")
        chainPath := filepath.Join(targetPath, "chain.pem")

        err = certificate.WriteCertificate(certPath, renewedCertificates.Certificate)
        if err != nil {
            return nil, err
        }

        err = certificate.WritePrivateKey(keyPath, renewedCertificates.PrivateKey)
        if err != nil {
            return nil, err
        }

        err = certificate.WriteCertificate(chainPath, renewedCertificates.IssuerCertificate)
        if err != nil {
            return nil, err
        }

        return renewedCertificates, nil
    }

    return certificates, nil
}

func RenewCertificatePeriodically(client *lego.Client, certificates *certificate.Resource, certOptions CertificateOptions) error {
	for {
		_, err := RenewCertificate(client, certificates, certOptions.RenewIntervalDays)
		if err != nil {
			return err
		}

		time.Sleep(time.Hour * 24 * time.Duration(certOptions.RenewIntervalDays))
	}
}