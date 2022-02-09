package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"os"
	"time"
)

func GenerateFiles(cn string, alternateDNS []string) error {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}

	keyBytes := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})

	caBytes, err := generateCA(cn, 10*365, key)
	if err != nil {
		return err
	}

	certBytes, err := generateSelfSignedCertificate(cn, alternateDNS, 10*365, key)
	if err != nil {
		return err
	}

	// write files
	file, err := os.Create("server.key")
	if err != nil {
		return err
	}
	file.Write(keyBytes)
	file.Close()

	file, err = os.Create("ca.crt")
	if err != nil {
		return err
	}
	file.Write(caBytes)
	file.Close()

	file, err = os.Create("server.crt")
	if err != nil {
		return err
	}
	file.Write(certBytes)
	file.Close()

	return nil
}

func generateCA(cn string, daysValid int, priv *rsa.PrivateKey) ([]byte, error) {
	template, err := getBaseCertTemplate(cn, []string{}, daysValid)
	if err != nil {
		return nil, err
	}

	// Override KeyUsage and IsCA
	template.KeyUsage = x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign
	template.IsCA = true

	return getCertificate(template, priv, template, priv)
}

func generateSelfSignedCertificate(cn string, alternateDNS []string, daysValid int, priv *rsa.PrivateKey) ([]byte, error) {
	template, err := getBaseCertTemplate(cn, alternateDNS, daysValid)
	if err != nil {
		return nil, err
	}

	return getCertificate(template, priv, template, priv)
}

func getCertificate(template *x509.Certificate, signeeKey *rsa.PrivateKey, parent *x509.Certificate, signingKey *rsa.PrivateKey) ([]byte, error) {
	derBytes, err := x509.CreateCertificate(rand.Reader, template, parent, signeeKey.Public(), signingKey)
	if err != nil {
		return nil, err
	}

	block := &pem.Block{Type: "CERTIFICATE", Bytes: derBytes}
	return pem.EncodeToMemory(block), nil
}

func getBaseCertTemplate(cn string, alternateDNS []string, daysValid int) (*x509.Certificate, error) {
	serialNumberUpperBound := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberUpperBound)
	if err != nil {
		return nil, err
	}

	return &x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			CommonName: cn,
		},
		DNSNames:  alternateDNS,
		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(time.Hour * 24 * time.Duration(daysValid)),
		KeyUsage:  x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{
			x509.ExtKeyUsageServerAuth,
			x509.ExtKeyUsageClientAuth,
		},
		BasicConstraintsValid: true,
	}, nil
}

func main() {
	GenerateFiles("dashboard", []string{"dashboard.contrail", "dashboard", "localhost", "127.0.0.1"})
}
