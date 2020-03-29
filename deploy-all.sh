#!/bin/bash
#bash deploy-consul-prometheus-grafana.sh
#sleep 60
bash join.sh
bash deploy-flagger.sh
bash deploy-ingress.sh
bash deploy-prefix-router.sh
bash deploy-svc.sh
