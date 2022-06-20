package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const certPath = "/var/opt/certman"

func main() {
	log.Println("uid:", os.Getuid())
	args := os.Args
	if len(args) < 2 {
		log.Fatal("need a command (init or a cert name followed by comma separated hostnames)")
	}
	cmd := args[1]
	if cmd == "init" {
		initCA()
		log.Println("initialized CA")
		return
	}

	name := cmd
	hosts := []string{"dev.rhyvu.com", "accounts.dev.rhyvu.com"}
	if len(args) > 2 {
		hosts = append(hosts, strings.Split(args[2], ",")...)
	}
	initDevCert(name, hosts)
}

func initDevCert(name string, hosts []string) {
	certName := filepath.Join(certPath, fmt.Sprintf("%s.crt", name))
	keyName := filepath.Join(certPath, fmt.Sprintf("%s.key", name))
	exists, err := certExists(fmt.Sprintf("%s.key", name))
	if err != nil {
		panic(err)
	}
	if exists {
		log.Fatalf("%s already exists", name)
	}
	cert := &x509.Certificate{
		// a serial should be unique within this CA, so a microsecond resolution of a unix timestamp is a reasonable
		// approach for dev work, though machine resolution might cause some duplicates
		SerialNumber: big.NewInt(time.Now().UnixMicro()),
		Subject: pkix.Name{
			// use a bogus address for the CA Subject information
			Organization:  []string{"Home Office"},
			Country:       []string{"US"},
			Province:      []string{"TN"},
			Locality:      []string{"Chattanooga"},
			StreetAddress: []string{"Trewhitt Street"},
			PostalCode:    []string{"37405"},
		},
		NotBefore: time.Now(),
		// a public issued cert is only accepted in a browser for up to 398 days
		NotAfter:     time.Now().AddDate(1, 0, 0),
		IPAddresses:  []net.IP{net.IPv4(127, 0, 0, 1), net.IPv6loopback},
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:     x509.KeyUsageDigitalSignature,
		DNSNames:     hosts,
		SubjectKeyId: []byte{1, 2, 3, 4, 6},
	}
	certPrivKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		panic(err)
	}

	caPrivKey, ca := getCA()
	certBytes, err := x509.CreateCertificate(rand.Reader, cert, ca, &certPrivKey.PublicKey, caPrivKey)
	if err != nil {
		panic(err)
	}
	certPEM := new(bytes.Buffer)
	err = pem.Encode(certPEM, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certBytes,
	})
	if err != nil {
		panic(err)
	}

	certPrivKeyPEM := new(bytes.Buffer)
	err = pem.Encode(certPrivKeyPEM, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(certPrivKey),
	})
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(keyName, certPrivKeyPEM.Bytes(), 0600)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(certName, certPEM.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}

func certExists(name string) (bool, error) {
	filename := filepath.Join(certPath, name)
	_, err := os.Stat(filename)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}

func getCA() (*rsa.PrivateKey, *x509.Certificate) {
	keyBytes, err := ioutil.ReadFile(filepath.Join(certPath, "ca.key"))
	if err != nil {
		panic(err)
	}
	keyPemBlock, _ := pem.Decode(keyBytes)
	key, err := x509.ParsePKCS1PrivateKey(keyPemBlock.Bytes)

	certBytes, err := ioutil.ReadFile(filepath.Join(certPath, "ca.crt"))
	if err != nil {
		panic(err)
	}
	certPemBlock, _ := pem.Decode(certBytes)
	cert, err := x509.ParseCertificate(certPemBlock.Bytes)

	return key, cert
}

func initCA() {
	exists, err := certExists("ca.key")
	if err != nil {
		panic(err)
	}
	if exists {
		panic("ca already exists")
	}
	ca := &x509.Certificate{
		// a serial should be unique within this CA, so a microsecond resolution of a unix timestamp is a reasonable
		// approach for dev work, though machine resolution might cause some duplicates
		SerialNumber: big.NewInt(time.Now().UnixMicro()),
		Subject: pkix.Name{
			// use a bogus address for the CA Subject information
			Organization:  []string{"Home Office"},
			Country:       []string{"US"},
			Province:      []string{"TN"},
			Locality:      []string{"Chattanooga"},
			StreetAddress: []string{"Trewhitt Street"},
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
	caPrivKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		panic(err)
	}
	caBytes, err := x509.CreateCertificate(rand.Reader, ca, ca, &caPrivKey.PublicKey, caPrivKey)
	if err != nil {
		panic(err)
	}
	caPEM := new(bytes.Buffer)
	err = pem.Encode(caPEM, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: caBytes,
	})
	if err != nil {
		panic(err)
	}
	caPrivKeyPEM := new(bytes.Buffer)
	err = pem.Encode(caPrivKeyPEM, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(caPrivKey),
	})
	if err != nil {
		panic(err)
	}

	keyName := filepath.Join(certPath, "ca.key")
	err = ioutil.WriteFile(keyName, caPrivKeyPEM.Bytes(), 0600)
	if err != nil {
		panic(err)
	}
	certName := filepath.Join(certPath, "ca.crt")
	err = ioutil.WriteFile(certName, caPEM.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}
