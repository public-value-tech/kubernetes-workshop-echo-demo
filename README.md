# kubernetes-workshop-echo-demo

deploy a demo application with built in config and with port 8081

## How to deploy

### Automatically with sed

Execute

```bash
sed -i "s/REPLACEME/$(whoami | sed 's/_/-/g')/g" kustomization.yml ingress.yml
```

### Manually

Search for all occurences of `REPLACEME`. Replace it with your account name, e.g. schulung-5.

### Apply 

Execute the following command to deploy the demo application to your namespace. 

```bash

kubectl apply -k .

```

### Checking health

If there are no errors during apply, check the output of `kubectl get ingress` to verify that the Ingress has been created.

Now you can browse to the shown address.

## Build with docker

```bash
docker build app -t europe-west3-docker.pkg.dev/swr-schulung-basic-1/swr-schulung-basic-1/echo-demo:latest
docker push europe-west3-docker.pkg.dev/swr-schulung-basic-1/swr-schulung-basic-1/echo-demo:latest
```

## Build with buildah

```bash
buildah bud --tag europe-west3-docker.pkg.dev/swr-schulung-basic-1/swr-schulung-basic-1/echo-demo:latest
buildah push europe-west3-docker.pkg.dev/swr-schulung-basic-1/swr-schulung-basic-1/echo-demo:latest
```
