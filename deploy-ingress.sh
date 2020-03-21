kubectl config use-context kind-dc1
helm upgrade -i ingress ingress-chart --set name=ingress --set controller.tag=$IMAGE_ID --set dc=dc1

kubectl config use-context kind-dc2
helm upgrade -i ingress ingress-chart --set name=ingress --set controller.tag=$IMAGE_ID --set dc=dc2
