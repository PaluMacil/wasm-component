package main

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"log"
	"math/big"
	"os"
	"time"
)

func main() {
	log.Println("uid:", os.Getuid())
}

func ca() *x509.Certificate {
	return &x509.Certificate{
		// a serial should be unique within this CA, so a microsecond resolution of a unix timestamp is a reasonable
		// approach for dev work, though machine resolution might cause some duplicates
		SerialNumber: big.NewInt(time.Now().UnixMicro()),
		Subject: pkix.Name{
			// use a bogus address for the CA Subject information
			Organization:  []string{"Home Office"},
			Country:       []string{"US"},
			Province:      []string{"TN"},
			Locality:      []string{"Chattanooga"},
			StreetAddress: []string{"17 Trewhitt Street"},
			PostalCode:    []string{"37405"},
		},
		NotBefore: time.Now(),
		// a dev CA can safely be valid for a long time, so let's give it 10 years
		NotAfter: time.Now().AddDate(10, 0, 0),
		// mark this as a CA
		IsCA: true,

		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}
}
