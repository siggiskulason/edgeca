# edgeca
**EdgeCA** is an ephemeral, in-memory CA providing service mesh machine identities.
It automates the management and issuance of TLS certificates. It can either run with a self-certificated Root CA certificate or use an issuing certificate retrieved using the [Venefi vCert](https://github.com/Venafi/vcert) software.

It solves the many limitations of the embedded service mesh CAs by providing developers a fast, easy, and integrated source of machine identities whilst also providing security teams with the required policy and oversight.  

It also enables ephemeral certificate-based authorization, which reduces the need for permanent access credentials, explicit access revocation or traditional SSH key management. 

EdgeCA is currently at version 0.4.0 and is meant for evaluation only. 

## Compiling edgeca

The first step is to compile and install EdgeCA:

```
git clone https://github.com/edgeca-org/edgeca.git
cd edgeca
go install ./cmd/edgeca
go install ./cmd/edgecad
```

## Start up **edgecad**
EdgeCA consists of two applications:
- edgeca is the command line interface (CLI) application you will use to create CSRs and certificates
- edgecad is a server which edgeca connects to. It is the core Ephemeral CA engine and signs the certificates. It can run in three different modes.

To get help, type
```
edgecad -h
```

The output will show:

```
EdgeCA can run in three modes
	
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
	./edgecad --url "..." --zone "..." --token "..." 
	
	EdgeCA gets an issuing certificate using the TPP token.
	It reads in the policy and default configuration from the TPP server 

Note: In all three modes, the server writes certificates to the location
specified by "tls-certs". These certificates are required by the edgeca client
for encryption and authentication.

Usage:
  edgecad [flags]

Flags:
  -c, --ca-cert string     Issuing Certificate File
  -k, --ca-key string      Issuing Certificate Key File
  -h, --help               help for edgecad
  -p, --policy string      Policy File
      --port int           Port number to use for this server (default 50025)
  -d, --tls-certs string   Directory to write gRPC TLS Client certificates to (default "/home/sidar/.edgeca/certs")
  -t, --token string       TPP Token
  -u, --url string         TPP URL
  -z, --zone string        TPP Zone

```

### Mode 1: Self-signed

Start by running edgeca in a terminal window. The simplest setup is to use a self-signed certificate. All you need then to do is to type

```
edgecad
```

EdgeCA will then start up 
```
EdgeCA daemon 0.3.0 starting up
edgeCA: 2021/01/19 16:24:16 Mode 1 (Using self-signed issuing certificate and key)
edgeCA: 2021/01/19 16:24:16 Generating self signed Root CA Certificate
edgeCA: 2021/01/19 16:24:16 Generating self signed Sub CA Certificate
edgeCA: 2021/01/19 16:24:16 Creating TLS server certificate for server
edgeCA: 2021/01/19 16:24:17 Writing TLS Client certificate to /home/me/.edgeca/certs/edgeca-client-cert.pem
edgeCA: 2021/01/19 16:24:17 Writing TLS Client key to /home/me/.edgeca/certs/edgeca-client-key.pem
edgeCA: 2021/01/19 16:24:17 Writing Root CA Certificate to  /home/me/.edgeca/certs/CA.pem
edgeCA: 2021/01/19 16:24:17 Starting gRPC CA server on port 50025
```

There are some optional parameters which you can use
- **--port** is used to specify the TCP/IP port to listen on. By defult EdgeCA uses 50025
- **--tls-certs** is used to specify where to write the certificates which are generated by EdgeCA for the TLS connection between EdgeCA server and client. These certificates need to be accessible by the client. By default both server and client use ~/.edgeca/certs but if you run the server and client on two different devices, then the certificates need to be copied to where they can be accessed by the client
- **--policy** can be used to Specify an [Open Policy Agent](https://www.openpolicyagent.org/) policy file. The policy should be in the "edgeca" package and be named "csr_policy". A sample file would be:

### (Optional) Using policies

Example file:

```
package edgeca

	csr_policy {

        re_match(`^ACME inc$`, input.csr.Subject.Organization[0])
 
    }

```

EdgeCA will transform Venafi TPP policies into OPA policies so if you are using TPP and have a policy defined, the edgecad will print out the OPA policy it generated. 

To use a manually generated policy file, specify it using the **-p** parameter 

```
edgecad -p opa/opa.rego
```

### Mode 3: Use TPP
In this mode, you need to provide a TPP access token when starting up edgecad. 

To use, specify the TPP URL, Zone and Access token. EdgeCA will then get an issuing certificate and policy information from the TPP server. 

```
./edgecad --url "..." --zone "..." --token "..." 
```

## Use the EdgeCA CLI

Once the server is up and running, you can use the edgeca CLI to issue commands.
```
edgeca -h
EdgeCA is an ephemeral certificate authority

Usage:
  edgeca [command]

Available Commands:
  gencert     Create a Certificate
  gencsr      Create a CSR
  help        Help about any command

Flags:
  -h, --help   help for edgeca

Use "edgeca [command] --help" for more information about a command.
```

It supports two commands
- **gencsr** generates a CSR (Certificate Signing Request) used 
- **gencert** returns a signed certificate 

### gencsr  
```
gencsr -h
Create a CSR

Usage:
  edgeca gencsr [flags]

Flags:
      --cn string             Common Name (required)
  -c, --country string        Country
      --csr string            Output file for CSR
  -h, --help                  help for gencsr
      --key string            Output file for private key used to sign CSR
  -l, --locality string       Locality
  -o, --organization string   Organization
      --ou string             Organizational Unit
      --port int              TLS port of gRPC server (default 50025)
  -r, --refresh               Refresh the list of default values from the current policy
  -s, --server string         EdgeCA gRPC server name 
      --st string             State/Province
  -d, --tls-certs string      Location of certs for gRPC authentication 
```
To create a simple CSR file, do:

```
edgeca gencsr --cn localhost --organization ACME --csr test.csr --key test.key
```

If you are running edgecad in TPP mode, then it might be using policies with default values for CSRs. running edgeca with the **-r** argument will cause it to connect to edgecad, using gRPC and edgecad then use TPP to refresh information about the policy.

If for instance, the TPP policy specifies a default Organization of ABC, then the CSR generation will output:

``` 
edgeca gencsr --cn localhost --csr test.csr --key test.key -r
2021/01/19 17:03:30 Connecting to edgeca server to refresh policy information
2021/01/19 17:03:30 Loading TLS certificates from /home/me/.edgeca/certs
2021/01/19 17:03:30 Updated configuration file /home/me/.edgeca/configuration.yaml
2021/01/19 17:03:31 Generated CSR for [ CN=localhost,O=ABC ]
``` 

The default values are written into a configuration file so if we generate another CSR, without requesting the refresh step, then the default organization value is used but no connection is made to the gRPC server or TPP server:
```
edgeca gencsr --cn localhost --csr test.csr --key test.key
2021/01/19 17:05:55 Generated CSR for [ CN=localhost,O=ABC ]
```

### gencsr  
gencert is used to sign a CSR and return a signed certificate. If any values in the CSR don't match required values in the OPA policy, then it is rejected and the certificate not signed.

Generate a Certificate. Optional filenames can be specified for the certificate and private key
```
edgeca gencert -i test.csr -o test.cert -k test.key
```

A policy setting may reject the CSR
```
2020/11/26 22:32:45 could not get: rpc error: code = Unknown desc = Policy rejected CSR
```

otherwise the certificate and key are output

```
edgeca gencert -i test.csr -o test.cert -k test.key
2021/01/19 17:09:47 Connecting to edgeca server to sign certificate
2021/01/19 17:09:47 Loading TLS certificates from /home/me/.edgeca/certs
2021/01/19 17:09:47 Wrote Certificate to test.cert
2021/01/19 17:09:47 Wrote key to test.key
```
