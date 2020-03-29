helm delete --kube-context=kind-dc1 consul
helm delete --kube-context=kind-dc2 consul

kubectl config use-context kind-dc1
kubectl delete pvc data-default-consul-server-0
kubectl delete pvc data-default-consul-server-1
kubectl delete pvc data-default-consul-server-2

kubectl config use-context kind-dc2
kubectl delete pvc data-default-consul-server-0
kubectl delete pvc data-default-consul-server-1
kubectl delete pvc data-default-consul-server-2

