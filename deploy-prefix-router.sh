function build() {
  pushd prefixrouter
  make
  popd
}

#build


IMAGE_ID=$(docker images oleksiyp/prefixrouter:latest --format "{{.ID}}")

docker tag oleksiyp/prefixrouter:latest oleksiyp/prefixrouter:$IMAGE_ID

kind load docker-image oleksiyp/prefixrouter:$IMAGE_ID --name dc1
kind load docker-image oleksiyp/prefixrouter:$IMAGE_ID --name dc2

kubectl config use-context kind-dc1
kubectl apply -f prefixrouter/artifacts/prefixrouter/crd.yaml
helm upgrade -i prefixrouter prefixrouter-chart --set name=prefixrouter --set controller.tag=$IMAGE_ID --set dc=dc1

kubectl config use-context kind-dc2
kubectl apply -f prefixrouter/artifacts/prefixrouter/crd.yaml
helm upgrade -i prefixrouter prefixrouter-chart --set name=prefixrouter --set controller.tag=$IMAGE_ID --set dc=dc2
