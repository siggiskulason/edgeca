package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/edgeca-org/edgeca"
	"github.com/edgeca-org/edgeca/internal/daemon"
	"github.com/edgeca-org/edgeca/internal/daemon/config"
	"github.com/edgeca-org/edgeca/internal/daemon/policies"
	"github.com/edgeca-org/edgeca/internal/daemon/state"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "edgecad",
	Short: "EdgeCA is an ephemeral certificate authority",
	Long: `EdgeCA can run in three modes
	
	Mode 1: Self-signed
	-------------------
	./edgecad [-p policy_file]
	
	In this mode, EdgeCA starts up, creates a self-signed certificate 
	and optionally reads in an OPA policy file. 
	
	Mode 2:  Bring your own CA Certificate
	--------------------------
	./edgecad [-p policy_file] -c certificate.pem -k key.pem
	
	In this mode, EdgeCA starts up, reads the CA certificate and key
	from the provided PEM files and optionally reads in an OPA policy file 
		
	Mode 3: Use TPP
	-------------------------------
	./edgecad -t TPP-token
	
	EdgeCA gets an issuing certificate using the TPP token.
	It reads in the policy and default configuration from the TPP server 

Note: In all three modes, the server writes certificates to the location
specified by "tls-certs". These certificates are required by the edgeca client
for encryption and authentication.
	`,
	Run: func(cmd *cobra.Command, args []string) {

		grpcMain()

	}}

var policy, defaultConfig, tppToken, tppURL, tppZone, caCert, caKey, tlsCertDir string
var tlsPort int

func init() {

	rootCmd.Flags().StringVarP(&policy, "policy", "p", "", "Policy File")

	rootCmd.Flags().StringVarP(&caCert, "ca-cert", "c", "", "Issuing Certificate File")
	rootCmd.Flags().StringVarP(&caKey, "ca-key", "k", "", "Issuing Certificate Key File")

	rootCmd.Flags().StringVarP(&tppToken, "token", "t", "", "TPP Token")
	rootCmd.Flags().StringVarP(&tppURL, "url", "u", "", "TPP URL")
	rootCmd.Flags().StringVarP(&tppZone, "zone", "z", "", "TPP Zone")

	tlsCertDir = config.GetDefaultTLSCertDir()
	rootCmd.Flags().StringVarP(&tlsCertDir, "tls-certs", "d", tlsCertDir, "Directory to write gRPC TLS Client certificates to")

	tlsPort = config.GetDefaultTLSPort()
	rootCmd.Flags().IntVarP(&tlsPort, "port", "", tlsPort, "Port number to use for this server")

}

// Execute the commands
func grpcMain() {
	fmt.Println("EdgeCA daemon " + edgeca.Version + " starting up")
	log.SetPrefix("edgeCA: ")

	if tppToken != "" || tppURL != "" || tppZone != "" {
		mode3UseTPP()
	} else if caCert != "" || caKey != "" {
		mode2BYOCert()

	} else {
		mode1SelfCert()
	}

	daemon.StartGrpcServer(tlsPort)

}

func mode1SelfCert() {

	if policy != "" {
		policies.LoadPolicy(policy)
	}

	log.Println("Mode 1 (Using self-signed issuing certificate and key)")
	state.InitState(tlsCertDir)

}

func mode2BYOCert() {
	if policy != "" {
		policies.LoadPolicy(policy)
	}

	log.Println("Mode 2 (Using provided issuing certificate and key).")
	err := state.InitStateUsingCerts(caCert, caKey, tlsCertDir)

	if err != nil {
		log.Fatalf("Error: %v", err.Error())
	}
}

func mode3UseTPP() {
	if caCert != "" || caKey != "" {
		log.Fatalln("Mode 3 (Using TPP). Error: If TPP-Token is specified, then CA-Cert and CA-Key can't also be specified. ")
	}
	if tppToken == "" || tppURL == "" || tppZone == "" {
		log.Fatalln("Mode 3 (Using TPP). Error: TPP Token, URL and Zone all need to be specified.")
	}

	if policy != "" || defaultConfig != "" {
		log.Println("Mode 3 (Using TPP). Warning: If TPP-Token is specified, policy file settings are ignored.")
	}

	log.Println("Mode 3 (Using TPP). Connecting using specified TPP token, URL and Zone")

	err := state.InitStateUsingTPP(tppURL, tppZone, tppToken, tlsCertDir)

	if err != nil {
		log.Fatalf("TPPLogin error: %v", err.Error())
	} else {
		log.Printf("TPPLogin OK")
	}
}

// Execute the commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
