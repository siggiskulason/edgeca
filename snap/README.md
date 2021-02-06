# EdgeCA
**EdgeCA** is an ephemeral, in-memory CA providing service mesh machine identities.

This early release is meant for evaluation only. It consists of two applications:
- **edgeca** is the command line interface (CLI) application you will use to create CSRs and certificates
- **edgecad** is a server which edgeca connects to. It is the core Ephemeral CA engine and signs the certificates

The two parts communicate in a secure way using gRPC.

To install the snap simply do

```
snap install edgeca
```

This is an early version of the snap. It will be updated for the server to run automatically and be accessible with "snap set" commands. However, at this stage, the two applications are installed as
- **edgeca** and
- **edgeca.d** 

To run EdgeCA, open a terminal window and run 

```
edgeca.d
```

With any choosen parameters as per the standard EdgeCA [instructions](../docs) and then run the client with

```
edgeca
```


