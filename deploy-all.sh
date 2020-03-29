#!/bin/bash
bash deploy-dashboard.sh
bash deploy-vault.sh
bash deploy-consul-k8s.sh
bash deploy-consul.sh
bash join.sh
bash deploy-monitoring.sh
bash deploy-flagger.sh
bash deploy-ingress.sh
bash deploy-prefix-router.sh
bash build-svc.sh
bash deploy-svc.sh
bash ingress-js/build.sh
