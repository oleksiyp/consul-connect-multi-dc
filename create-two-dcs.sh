

kind create cluster --name "dc1" --config cluster.yaml
kind create cluster --name "dc2" --config cluster.yaml

function setup() {
  kubectl apply -f dashboard.yaml

#  helm upgrade -i -f prometheus-values.yaml prometheus prometheus
#  helm upgrade -i -f grafana-values.yaml grafana stable/grafana

  kubectl label nodes $1-worker consul=server --overwrite
  kubectl label nodes $1-worker2 consul=server --overwrite
  kubectl label nodes $1-worker3 consul=server --overwrite
  kubectl label nodes $1-worker4 consul=client --overwrite
  kubectl label nodes $1-worker5 consul=client --overwrite

  helm upgrade -i consul ./consul-helm --values consul-values.yaml --set global.datacenter=$1


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
setup dc1

kubectl config use-context kind-dc2
setup dc2
