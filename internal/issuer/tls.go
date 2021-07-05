/*******************************************************************************
 * Copyright 2021 EdgeSec OÜ
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *
 *******************************************************************************/

package issuer

import (
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"errors"
	"io/ioutil"
	"log"
)

func GenerateTLSServerCert(server string, parentCert *x509.Certificate, parentKey *rsa.PrivateKey) (*tls.Certificate, error) {
	log.Println("Creating TLS server certificate for ", server)
	subject := pkix.Name{
		Organization:       []string{"EdgeCA"},
		OrganizationalUnit: []string{},
		CommonName:         server,
		Locality:           []string{},
		Province:           []string{},
		Country:            []string{},
	}

	pemCert, pemKey, _, err := GeneratePemCertificate(subject, parentCert, parentKey)
	if err != nil {
		return nil, err
	}

	cert, err := tls.X509KeyPair(pemCert, pemKey)
	return &cert, err
}

func LoadCAServerCert(filename string) (*x509.CertPool, error) {

	pemCert, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemCert) {
		return nil, errors.New("Could not append CA Certificate")
	}

	return certPool, nil
}

func GenerateTLSClientCert(server string, parentCert *x509.Certificate, parentKey *rsa.PrivateKey, certfilename string, keyfilename string) (*tls.Certificate, error) {

	subject := pkix.Name{
		Organization:       []string{"EdgeCA"},
		OrganizationalUnit: []string{},
		CommonName:         server,
		Locality:           []string{},
		Province:           []string{},
		Country:            []string{},
	}

	pemCert, pemKey, _, err := GeneratePemCertificate(subject, parentCert, parentKey)
	if err != nil {
		return nil, err
	}

	cert, err := tls.X509KeyPair(pemCert, pemKey)

	if certfilename != "" {
		err := ioutil.WriteFile(certfilename, pemCert, 0644)
		if err != nil {
			log.Fatalf("Error writing output to %s: %v", certfilename, err)
		}
		log.Printf("Writing TLS Client certificate to %s", certfilename)
	}
	if keyfilename != "" {
		err := ioutil.WriteFile(keyfilename, pemKey, 0644)
		if err != nil {
			log.Fatalf("Error writing output to %s: %v", keyfilename, err)
		}
		log.Printf("Writing TLS Client key to %s", keyfilename)
	}
	return &cert, err
}
