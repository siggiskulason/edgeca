echo "Copying gRPC client certificates to a shared custom config volume"
cp ~/.edgeca/certs/edgeca-client-cert.pem ./custom-config/clientcert.pem
cp ~/.edgeca/certs/edgeca-client-key.pem ./custom-config/clientkey.pem
chmod -R a+rw ./custom-config

echo "Start up Envoy with a custom configuration"
sudo docker run --rm -it \
      -v $(pwd)/custom-config:/custom-config\
      -p 9901:9901 \
      -p 10000:10000 \
      envoyproxy/envoy-dev:747944b30b5556b07a5bffdea46fcea89404b9f4 \
        -c /custom-config/edgeca-envoy.yaml
