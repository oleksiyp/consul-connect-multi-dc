helm delete --kube-context=kind-dc1 vault

kubectl config use-context kind-dc1
kubectl delete pvc data-vault-0
