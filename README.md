# edgeca
Ephemeral Certificate Authority

## Getting started:

### get the source repository

```
go get -u github.com/edgeca-org/edgeca
```

### install edgeca and edgecad

```
go install github.com/edgeca-org/edgeca/cmd/edgeca
go install github.com/edgeca-org/edgeca/cmd/edgecad
```

### Start up the daemon
- An optional policy file can be specified. This sample file contains a simple rule requiring the issuing Organization to be "Venafi"

```
edgecad -p opa/opa.rego
```

### In a different window, enter client commands

Get the self-signed root certificate. Specify an optional output parameter to write it to a file - otherwise it is output to the console
```
edgeca getroot -o ca.crt
```

Review the current policy file, if one was specified
```
edgeca getpolicy
```

Generate a CSR
```
./edgeca gencsr --cn localhost --organization ACME --csr test.csr --key test.key
```

Generate a Certificate
```
edgeca gencert -i test.csr
```

Note that the policy rejects the CSR
```
2020/11/26 22:32:45 could not get: rpc error: code = Unknown desc = Policy rejected CSR
```

Regenerate the CSR
```
edgeca gencsr --cn localhost --organization Venafi --csr test.csr --key test.key
```

and regenerate the certificate. Optional filenames can be specified for the certificate and private key
```
edgeca gencert -i test.csr -o test.cert -k test.key
```

