#!/bin/bash

cd $(dirname $0)

kubectl config use-context kind-dc1
DC1_IPS=$(kubectl get nodes -l '!node-role.kubernetes.io/master' -o jsonpath={.items.*.status.addresses[0].address})
kubectl config use-context kind-dc2
DC2_IPS=$(kubectl get nodes -l '!node-role.kubernetes.io/master' -o jsonpath={.items.*.status.addresses[0].address})

cat >default.conf <<EOF
upstream traffic_backend {
  random;
EOF
for IP in $DC1_IPS $DC2_IPS; do
  echo "    server $IP:32203;" >> default.conf
done
cat >>default.conf <<EOF
}

server {
    listen       80;

    location / {
        root   /usr/share/nginx/html;
        index  index.html index.htm;
    }

    error_page  404              /404.html;
    error_page   500 502 503 504  /50x.html;
    location = /50x.html {
        root   /usr/share/nginx/html;
    }

    location /traffic {
        proxy_pass http://traffic_backend;
        proxy_http_version 1.1;
    }
}
EOF

npm run build
docker build -t oleksiyp/ingress-js:latest .
