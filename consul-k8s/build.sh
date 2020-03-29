make dev-docker

IMAGE_ID=$(docker images consul-k8s-dev:latest --format "{{.ID}}")

docker tag consul-k8s-dev:latest consul-k8s-dev:$IMAGE_ID

kind load docker-image consul-k8s-dev:$IMAGE_ID --name dc1
kind load docker-image consul-k8s-dev:$IMAGE_ID --name dc2

