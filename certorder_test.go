package certorder

import (
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"
)

// Define a test config path for your config.json
const testConfigPath = "path/to/your/test/config.json"

func TestCreateCertificate(t *testing.T) {
	// Load test certificate options from the test config file
	certOptions, err := LoadCertificateOptions(testConfigPath)
	if err != nil {
		t.Errorf("LoadCertificateOptions failed: %v", err)
		return
	}

	// Create a temporary test directory for certificates
	testDir := "path/to/temporary/test/directory"
	defer os.RemoveAll(testDir)

	// Override the CertTargetPath with the test directory
	certOptions.CertTargetPath = testDir

	// Call the CreateCertificate function
	certificates, err := CreateCertificate(certOptions)
	if err != nil {
		t.Errorf("CreateCertificate failed: %v", err)
		return
	}

	// Verify the created certificates
	certPath := filepath.Join(testDir, "cert.pem")
	keyPath := filepath.Join(testDir, "key.pem")
	chainPath := filepath.Join(testDir, "chain.pem")

	if _, err := os.Stat(certPath); os.IsNotExist(err) {
		t.Errorf("Certificate file not found: %v", err)
	}

	if _, err := os.Stat(keyPath); os.IsNotExist(err) {
		t.Errorf("Key file not found: %v", err)
	}

	if _, err := os.Stat(chainPath); os.IsNotExist(err) {
		t.Errorf("Chain file not found: %v", err)
	}
}

func TestRenewCertificate(t *testing.T) {
	// Load test certificate options from the test config file
	certOptions, err := LoadCertificateOptions(testConfigPath)
	if err != nil {
		t.Errorf("LoadCertificateOptions failed: %v", err)
		return
	}

	// Create a temporary test directory for certificates
	testDir := "path/to/temporary/test/directory"
	defer os.RemoveAll(testDir)

	// Override the CertTargetPath with the test directory
	certOptions.CertTargetPath = testDir

	// Create a temporary test client
	client, err := CreateLegoClient(certOptions)
	if err != nil {
		t.Errorf("CreateLegoClient failed: %v", err)
		return
	}

	// Create a temporary test certificate
	initialCertificates, err := CreateCertificate(certOptions)
	if err != nil {
		t.Errorf("CreateCertificate failed: %v", err)
		return
	}

	// Call the RenewCertificate function
	renewedCertificates, err := RenewCertificate(client, initialCertificates, certOptions.RenewIntervalDays, certOptions)
	if err != nil {
		t.Errorf("RenewCertificate failed: %v", err)
		return
	}

	// Verify the renewed certificates
	certPath := filepath.Join(testDir, "cert.pem")
	keyPath := filepath.Join(testDir, "key.pem")
	chainPath := filepath.Join(testDir, "chain.pem")

	if _, err := os.Stat(certPath); os.IsNotExist(err) {
		t.Errorf("Renewed certificate file not found: %v", err)
	}

	if _, err := os.Stat(keyPath); os.IsNotExist(err) {
		t.Errorf("Renewed key file not found: %v", err)
	}

	if _, err := os.Stat(chainPath); os.IsNotExist(err) {
		t.Errorf("Renewed chain file not found: %v", err)
	}

	// Optionally, you can also validate the renewed certificate details
	x509Cert, _ := pem.Decode([]byte(certificates.Certificate))
	parsedCert, err := x509.ParseCertificate(x509Cert.Bytes)
	if err != nil {
		t.Errorf("Error parsing renewed certificate: %v", err)
	}

	// Add more assertions as needed for specific certificate details.
}
