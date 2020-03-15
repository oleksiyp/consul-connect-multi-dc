function build() {
  pushd ingress-router-controller
  mvn compile com.google.cloud.tools:jib-maven-plugin:1.8.0:dockerBuild -Dimage=ingress-router-controller:latest -Djib.from.image=openjdk:8-jdk-alpine || exit 1
  popd
}

#build

IMAGE_ID=$(docker images ingress-router-controller:latest --format "{{.ID}}")

docker tag ingress-router-controller:latest ingress-router-controller:$IMAGE_ID

kind load docker-image ingress-router-controller:$IMAGE_ID --name dc1
kind load docker-image ingress-router-controller:$IMAGE_ID --name dc2

kubectl config use-context kind-dc1
helm upgrade -i ingress-router ingress-router-chart --set name=ingress-router --set controller.tag=$IMAGE_ID --set dc=dc1

kubectl config use-context kind-dc2
helm upgrade -i ingress-router ingress-router-chart --set name=ingress-router --set controller.tag=$IMAGE_ID --set dc=dc2
