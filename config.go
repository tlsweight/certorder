package certorder

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// CertificateOptions represents the certificate configuration.
type CertificateOptions struct {
	CAProvider       string `json:"ca_provider"`
	Email            string `json:"email"`
	CommonName       string `json:"common_name"`
	OrgUnit          string `json:"org_unit"`
	State            string `json:"state"`
	CountryCode      string `json:"country_code"`
	KeyType          string `json:"key_type"`
	KeyBits          int    `json:"key_bits"`
	RenewCertificate  bool   `json:"renew_certificate"`
	DNSProvider      string `json:"dns_provider"`
	DNSAPIKey        string `json:"dns_api_key"`
	RenewIntervalDays int    `json:"renew_interval_days"`
	CertTargetPath    string `json:"cert_target_path"` // Add this field
}

// LoadCertificateOptions loads the certificate configuration from a JSON file.
func LoadCertificateOptions(filename string) (*CertificateOptions, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var certOptions CertificateOptions
	err = json.Unmarshal(data, &certOptions)
	if err != nil {
		return nil, err
	}

	// Use a default certificate target path if not specified
	if certOptions.CertTargetPath == "" {
		certOptions.CertTargetPath = "/tmp/certorder/"
	}

	return &certOptions, nil
}
