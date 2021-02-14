# edgeca
**EdgeCA** is an ephemeral, in-memory CA providing service mesh machine identities.
It automates the management and issuance of TLS certificates. It can either run with a self-certificated Root CA certificate or use an issuing certificate retrieved using the [Venafi vCert](https://github.com/Venafi/vcert) software.

It solves the many limitations of the embedded service mesh CAs by providing developers a fast, easy, and integrated source of machine identities whilst also providing security teams with the required policy and oversight.  

It also enables ephemeral certificate-based authorization, which reduces the need for permanent access credentials, explicit access revocation or traditional SSH key management. 

EdgeCA is currently at version 0.4.0 and is meant for evaluation only. 

The easiest way to install the application is to use [snaps](./snap)

```
snap install edgeca
```

EdgeCA is open source, written in Go, and licenced with the Apache 2.0 Licence

See these [instructions](docs) on how to install and run EdgeCA
