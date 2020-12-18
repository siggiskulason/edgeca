package certs

import (
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"log"
)

func PemToRSAPrivateKey(pemKey []byte) (rsaKey *rsa.PrivateKey, err error) {

	key, err := pemToPrivateKey(pemKey)
	if err == nil {
		rsaKey = key.(*rsa.PrivateKey)
	}
	return
}

func PemToCert(pemCert []byte) (cert *x509.Certificate, err error) {

	der, _ := pem.Decode(pemCert)
	if der == nil {
		err = errors.New("Could not decode PEM cert  ")
		return
	}

	derBytes := der.Bytes

	certificate, err := x509.ParseCertificate(derBytes)
	if err != nil {
		log.Println("ParseCertificate failed:", err)
	}
	return certificate, err

}

func pemToPrivateKey(pemKey []byte) (key crypto.PrivateKey, err error) {

	der, _ := pem.Decode(pemKey)
	if der == nil {
		err = errors.New("Could not decode PEM key  ")
		return
	}

	derBytes := der.Bytes

	key, err = x509.ParsePKCS1PrivateKey(derBytes)
	if err != nil {
		key, err = x509.ParsePKCS8PrivateKey(derBytes)
		if err != nil {
			key, err = x509.ParseECPrivateKey(derBytes)
		}
	}
	return
}
