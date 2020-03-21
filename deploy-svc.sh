function build() {
  pushd service
  make
  popd
}

build

IMAGE_ID=$(docker images oleksiyp/service:latest --format "{{.ID}}")

docker tag oleksiyp/service:latest oleksiyp/service:$IMAGE_ID

kind load docker-image oleksiyp/service:$IMAGE_ID --name dc1
kind load docker-image oleksiyp/service:$IMAGE_ID --name dc2

kubectl config use-context kind-dc1
helm upgrade -i svc1 service-chart --set name=svc1 --set image.tag=$IMAGE_ID --set dc=dc1
#helm upgrade -i svc2 service-chart --set name=svc2 --set image.tag=$IMAGE_ID --set dc=dc1
#helm upgrade -i svc3 service-chart --set name=svc3 --set image.tag=$IMAGE_ID --set dc=dc1
#helm upgrade -i svc4 service-chart --set name=svc4 --set image.tag=$IMAGE_ID --set dc=dc1

kubectl config use-context kind-dc2
helm upgrade -i svc1 service-chart --set name=svc1 --set image.tag=$IMAGE_ID --set dc=dc2
#helm upgrade -i svc2 service-chart --set name=svc2 --set image.tag=$IMAGE_ID --set dc=dc2
#helm upgrade -i svc3 service-chart --set name=svc3 --set image.tag=$IMAGE_ID --set dc=dc2
#helm upgrade -i svc4 service-chart --set name=svc4 --set image.tag=$IMAGE_ID --set dc=dc2

