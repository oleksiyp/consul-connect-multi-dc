
kubectl config use-context kind-dc1
helm upgrade -i -f prometheus-values.yaml prometheus stable/prometheus
helm upgrade -i -f grafana-values.yaml grafana stable/grafana

kubectl config use-context kind-dc2
helm upgrade -i -f prometheus-values.yaml prometheus stable/prometheus
helm upgrade -i -f grafana-values.yaml grafana stable/grafana
