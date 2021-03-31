# EdgeCA
**EdgeCA** is an ephemeral, in-memory CA providing service mesh machine identities.

This early release is meant for evaluation only. It consists of two applications:
- **edgeca** is the command line interface (CLI) application you will use to create CSRs and certificates
- **edgecad** is a background service which edgeca connects to. It is the core Ephemeral CA engine and signs the certificates

The two parts communicate in a secure way using gRPC.

To install the snap simply do

```
snap install edgeca
```

This is an early version of the snap. The server part (edgecad) starts up as a background daemon by default and
does not need to be manually launched.

edgecad starts up in the default self-signed mode. To use TPP, set the following:

```
$ sudo snap set edgeca tpp.token="your token" 
$ sudo snap set edgeca tpp.zone="your zone"
$ sudo snap set edgeca tpp.url="your tpp url"
```

once those three have been set, edgecad will establish a TPP connection

To view the logs do

```
snap logs -f edgeca.edgecad 
```

The policy can likewise be set with

```
$ sudo snap set edgeca policy="policy filename" 
```

The client is run using

```
edgeca
```


