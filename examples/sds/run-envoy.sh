echo "Copying gRPC client certificates to a shared custom config volume"
cp ~/.edgeca/certs/edgeca-client-cert.pem ./custom-config/clientcert.pem
cp ~/.edgeca/certs/edgeca-client-key.pem ./custom-config/clientkey.pem
chmod -R a+rw ./custom-config

echo "Start up Envoy with a custom configuration"
envoy -c /custom-config/edgeca-envoy.yaml
