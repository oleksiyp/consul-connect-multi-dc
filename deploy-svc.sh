

function build() {
  pushd ../service
  mvn compile com.google.cloud.tools:jib-maven-plugin:1.8.0:dockerBuild -Dimage=service:latest -Djib.from.image=openjdk:8-jdk-alpine || exit 1
  popd
}


function setup() {
  helm upgrade -i    svc1 service-chart \
     --set name=svc1 \
     --set image.tag=$IMAGE_ID \
     --set dc=$1 \
     --set color=$3

#  helm upgrade -i svc2 service-chart \
#     --set name=svc2 \
#     --set image.tag=$IMAGE_ID \
#     --set dc=$1 \
#     --set color=$4

#  helm upgrade -i svc3 service-chart \
#     --set name=svc3 \
#     --set image.tag=$IMAGE_ID \
#     --set dc=$1 \
#     --set color=$5

#  helm upgrade -i svc4 service-chart \
#     --set name=svc4 \
#     --set image.tag=$IMAGE_ID \
#     --set dc=$1 \
#     --set color=$6
}

build

IMAGE_ID=$(docker images service:latest --format "{{.ID}}")

docker tag service:latest service:$IMAGE_ID

kind load docker-image service:$IMAGE_ID --name dc1
kind load docker-image service:$IMAGE_ID --name dc2

kubectl config use-context kind-dc1
setup dc1 dc2 blue magenta brown cyan

kubectl config use-context kind-dc2
setup dc2 dc1 blue red blue orange

