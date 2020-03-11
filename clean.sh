
kubectl config use-context kind-dc1
helm delete consul
kubectl delete pvc data-default-consul-server-0
kubectl delete pvc data-default-consul-server-1
kubectl delete pvc data-default-consul-server-2

kubectl config use-context kind-dc2
helm delete consul
kubectl delete pvc data-default-consul-server-0
kubectl delete pvc data-default-consul-server-1
kubectl delete pvc data-default-consul-server-2
