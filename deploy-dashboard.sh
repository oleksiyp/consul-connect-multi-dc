kubectl config use-context kind-dc1
kubectl apply -f dashboard.yaml

kubectl config use-context kind-dc2
kubectl apply -f dashboard.yaml
