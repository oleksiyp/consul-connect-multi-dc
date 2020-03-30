## Startup instruction
```
Prerequisites: docker, kind, go, helm, java, maven, node, vault(command line)

$ bash create-two-dcs.sh
$ bash deploy-all.sh
$ docker run -p 3000:80 oleksiyp/ingress-js:latest

Open [http://localhost:3000](http://localhost:3000)
```

### Control commands

### Deploy

bash control-svc.sh deploy dc1 svc1

### Off

bash control-svc.sh off dc1 svc1

### On

bash control-svc.sh on dc1 svc1
