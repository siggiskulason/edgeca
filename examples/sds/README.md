## Envoy Secret Discovery Service (SDS) support

EdgeCA has POC support for SDS. This is currently under development and can be enabled by running with the "--sds" flag.

To run this example, do the following

1. Start EdgeCA with SDS mode enabled

```
./run-edgeca.sh
```

2. In a separate window, start up Envoy

You first need to have [installed Envoy](https://www.envoyproxy.io/docs/envoy/latest/start/install)

If you have installed it using docker, then run

```
./run-envoy-docker.sh
```

if you have installed is as a native application (for instance using brew on MacOS) use

```
./run-envoy.sh
```

This will start up Envoy with a the configuration from `custom-config/edgeca-envoy.yaml`. This configuration
- creates a listener on port 10000
- gets the TLS certificate for the listener using SDS from EdgeCA. To connect to EdgeCA it uses gRPC and the gRPC TLS client certificates to authenticate against EdgeCA
- sets up a proxy redirecting to https://edgeca.org



3. To test this, use curl

```
curl --cacert ~/.edgeca/certs/CA.pem https://localhost:10000
```

This curl command will use the CA cert from EdgeCA. It will connect to Envoy, which will in turn use TLS certificates provided to it by EdgeCA.

The expected output is the webpage at https://edgeca.org, which is the one returned by the sample proxy configuration
