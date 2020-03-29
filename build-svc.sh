function build() {
  pushd service-ktor
  mvn package
  docker build . -t service
  popd
}

build

IMAGE_ID=$(docker images service:latest --format "{{.ID}}")

docker tag service:latest service:$IMAGE_ID

kind load docker-image service:$IMAGE_ID --name dc1
kind load docker-image service:$IMAGE_ID --name dc2
