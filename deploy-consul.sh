
CONSUL_K8S_IMAGE_ID=$(docker images consul-k8s-dev:latest --format "{{.ID}}")

function setup() {
  kubectl label nodes $1-worker consul=server --overwrite
  kubectl label nodes $1-worker2 consul=server --overwrite
  kubectl label nodes $1-worker3 consul=server --overwrite
  kubectl label nodes $1-worker4 consul=client --overwrite
  kubectl label nodes $1-worker5 consul=client --overwrite

  # required to make deploy repatable and idempotent
  # one first run not needed
  VAULT_TOKEN=root VAULT_ADDR=http://$VAULT_IP:8200 vault delete sys/mounts/connect-ca-$1

  helm upgrade -i consul ./consul-helm --values consul-values.yaml \
      --set global.imageK8S=consul-k8s-dev:$CONSUL_K8S_IMAGE_ID \
      --set global.datacenter=$1 \
      --set server.connectCA.vault.intermediatePKIPath=/connect-ca-$1 \
      --set server.connectCA.vault.address=http://$VAULT_IP:8200 \
      --wait

  cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: ConfigMap
metadata:
  name: coredns
  namespace: kube-system
data:
  Corefile: |
    .:53 {
        errors
        health
        ready
        kubernetes cluster.local in-addr.arpa ip6.arpa {
           pods insecure
           fallthrough in-addr.arpa ip6.arpa
           ttl 30
        }
        prometheus :9153
        forward . /etc/resolv.conf
        cache 30
        loop
        reload
        loadbalance
    }
    consul {
      errors
      cache 30
      forward . $(kubectl get svc consul-dns -o jsonpath={.spec.clusterIP})
    }
EOF
}

kubectl config use-context kind-dc1
VAULT_IP=$(kubectl get pod vault-0 -o jsonpath={.status.hostIP})
setup dc1

kubectl config use-context kind-dc2
setup dc2
