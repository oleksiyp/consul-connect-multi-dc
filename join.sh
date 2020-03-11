

function ips() {
        kubectl config use-context kind-$1 > /dev/null
	kubectl get nodes --selector=kubernetes.io/role!=master --selector=consul=server -o jsonpath={.items[*].status.addresses[?\(@.type==\"InternalIP\"\)].address}
}

function join_all() {
  for ip1 in `ips $1`; do
    for ip2 in `ips $2`; do
      echo "Joining $ip1 to $ip2"
      CONSUL_HTTP_ADDR=$ip1:8500 consul join -wan $ip2
    done
  done
}

join_all dc1 dc2
join_all dc2 dc1
