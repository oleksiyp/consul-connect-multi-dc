
cd $(dirname $(realpath $0))

if [[ $# -ne 3 ]]; then
  echo "usage: ./contorl-svc.sh deploy dc1 svc3"
  echo "usage: ./contorl-svc.sh on dc1 svc3"
  echo "usage: ./contorl-svc.sh off dc1 svc3"
  exit 1
fi

CMD=$1
shift

if [ "$CMD" == "deploy" ]; then
  IMAGE_ID=$(docker images service:latest --format "{{.ID}}")
  kubectl config use-context kind-$1
  helm upgrade -i $2 service-chart --set name=$2 --set image.tag=$IMAGE_ID --set dc=$1
elif [[ "$CMD" == "on" || "$CMD" == "off" ]]; then
  kubectl config use-context kind-$1
  for SVC in `kubectl get pods -o name | grep $2`; do
    SVC=$(echo $SVC | sed "s#pod/##")
    kubectl exec $SVC -c service -- wget -O "$RANDOM-$RANDOM.txt" "http://localhost:8080/$CMD"
  done
fi
