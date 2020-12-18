package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/edgeca-org/edgeca/internal/certs"
	"github.com/edgeca-org/edgeca/internal/cli/config"
	internalgrpc "github.com/edgeca-org/edgeca/internal/grpc"

	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"

	"github.com/spf13/cobra"
)

var commonName, organization, organizationalUnit, locality, province, country, keyFileName, csrFileName, tlsCertDir, csrHostName string
var refresh bool
var csrTLSPort int

func init() {
	var gencsr = &cobra.Command{
		Use:   "gencsr",
		Short: "Create a CSR",
		Run: func(cmd *cobra.Command, args []string) {

			generateCSR()

		}}

	rootCmd.AddCommand(gencsr)

	gencsr.Flags().StringVarP(&commonName, "cn", "", "", "Common Name (required)")
	gencsr.MarkFlagRequired("cn")
	gencsr.Flags().StringVarP(&organization, "organization", "o", "", "Organization")
	gencsr.Flags().StringVarP(&organizationalUnit, "ou", "", "", "Organizational Unit")
	gencsr.Flags().StringVarP(&locality, "locality", "l", "", "Locality")
	gencsr.Flags().StringVarP(&province, "st", "", "", "State/Province")
	gencsr.Flags().StringVarP(&country, "country", "c", "", "Country")
	gencsr.Flags().StringVarP(&keyFileName, "key", "", "", "Output file for private key used to sign CSR")
	gencsr.Flags().StringVarP(&csrFileName, "csr", "", "", "Output file for CSR")
	gencsr.Flags().BoolVarP(&refresh, "refresh", "r", false, "Refresh the list of default values from the current policy")
	tlsCertDir = config.GetDefaultTLSCertDir()
	gencsr.Flags().StringVarP(&tlsCertDir, "tls-certs", "d", tlsCertDir, "Location of certs for gRPC authentication")
	csrHostName = config.GetDefaultTLSHost()
	gencsr.Flags().StringVarP(&csrHostName, "server", "s", csrHostName, "EdgeCA gRPC server name")
	csrTLSPort = config.GetDefaultTLSPort()
	gencsr.Flags().IntVarP(&csrTLSPort, "port", "", csrTLSPort, "TLS port of gRPC server")
}

func generateCSR() {

	// only use gRPC if --refresh is set
	if refresh {
		log.Println("Connecting to edgeca server at " + csrHostName + " to refresh policy information")

		conn, c := grpcConnect(tlsCertDir, csrHostName, csrTLSPort)
		defer conn.Close()
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		policyReply, err := c.RequestPolicy(ctx, &internalgrpc.PolicyRequest{})
		if err != nil {
			log.Fatalf("could not get: %v", err)
		}

		o := policyReply.GetDefaultOrganization()
		ou := policyReply.GetDefaultOrganizationalUnit()
		p := policyReply.GetDefaultProvince()
		l := policyReply.GetDefaultLocality()
		country := policyReply.GetDefaultCountry()
		config.SetCSRConfiguration(o, ou, country, p, l)
		config.WriteConfiguration()
	}

	// then generate the CSR

	csrBytes, keyBytes, err := generatePemCSR()

	if csrFileName != "" {
		err = ioutil.WriteFile(csrFileName, csrBytes, 0644)
		if err != nil {
			log.Fatalf("Error writing CSR to %s: %v", csrFileName, err)
		}
	} else {
		fmt.Println(string(csrBytes))
	}

	if keyFileName != "" {
		err = ioutil.WriteFile(keyFileName, keyBytes, 0644)
		if err != nil {
			log.Fatalf("Error writing key to %s: %v", keyFileName, err)
		}
	} else {
		fmt.Println(string(keyBytes))
	}

}

// GeneratePemCSR generates a certificate
func generatePemCSR() (csrBytes []byte, privatekeyBytes []byte, err error) {
	//	var oidEmailAddress = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 9, 1}

	//	emailAddress := "test@example.com"
	subj := pkix.Name{
		CommonName: commonName,
	}

	if organization == "" {
		organization = config.GetDefaultOrganization()
		if organization != "" {
			subj.Organization = []string{organization}
		}
	}

	if organizationalUnit == "" {
		organizationalUnit = config.GetDefaultOrganizationalUnit()
		if organizationalUnit != "" {
			subj.OrganizationalUnit = []string{organizationalUnit}
		}
	}

	if locality == "" {
		locality = config.GetDefaultLocality()
		if locality != "" {
			subj.Locality = []string{locality}
		}
	}

	if province == "" {
		province = config.GetDefaultProvince()
		if province != "" {
			subj.Province = []string{province}
		}
	}
	if country == "" {
		country = config.GetDefaultCountry()
		if country != "" {
			subj.Country = []string{country}
		}
	}

	privateRSAKey, err := certs.GenerateRSAKey()

	derBytes, err := certs.GenerateCSR(subj, privateRSAKey)

	csrBytes = pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE REQUEST", Bytes: derBytes})
	privatekeyBytes = pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateRSAKey)})
	log.Println("Generated CSR for [", subj, "]")
	return csrBytes, privatekeyBytes, err
}
