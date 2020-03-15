function build() {
  pushd flagger
  make
  popd
}

build

IMAGE_ID=$(docker images weaveworks/flagger:latest --format "{{.ID}}")

docker tag weaveworks/flagger:latest weaveworks/flagger:$IMAGE_ID

kind load docker-image weaveworks/flagger:$IMAGE_ID --name dc1
kind load docker-image weaveworks/flagger:$IMAGE_ID --name dc2

CONSUL_DC1=`docker inspect --format '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' dc1-server1`
CONSUL_DC2=`docker inspect --format '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' dc2-server1`

helm upgrade --kube-context kind-dc1 \
  -i flagger flagger/charts/flagger \
  --set image.tag=$IMAGE_ID \
  --set metricsServer=http://prometheus-server \
  --set meshProvider=connect \
  --set connect.address=$CONSUL_DC1:8500

helm upgrade --kube-context kind-dc2 \
  -i flagger flagger/charts/flagger \
  --set image.tag=$IMAGE_ID \
  --set metricsServer=http://prometheus-server \
  --set meshProvider=connect \
  --set connect.address=$CONSUL_DC2:8500
