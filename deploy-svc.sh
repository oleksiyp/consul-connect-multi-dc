IMAGE_ID=$(docker images service:latest --format "{{.ID}}")

kubectl config use-context kind-dc1
helm upgrade -i svc1 service-chart --set name=svc1 --set image.tag=$IMAGE_ID --set dc=dc1
helm upgrade -i svc2 service-chart --set name=svc2 --set image.tag=$IMAGE_ID --set dc=dc1
helm upgrade -i svc3 service-chart --set name=svc3 --set image.tag=$IMAGE_ID --set dc=dc1
helm upgrade -i svc4 service-chart --set name=svc4 --set image.tag=$IMAGE_ID --set dc=dc1

kubectl config use-context kind-dc2
helm upgrade -i svc1 service-chart --set name=svc1 --set image.tag=$IMAGE_ID --set dc=dc2
helm upgrade -i svc2 service-chart --set name=svc2 --set image.tag=$IMAGE_ID --set dc=dc2
helm upgrade -i svc3 service-chart --set name=svc3 --set image.tag=$IMAGE_ID --set dc=dc2
helm upgrade -i svc4 service-chart --set name=svc4 --set image.tag=$IMAGE_ID --set dc=dc2
