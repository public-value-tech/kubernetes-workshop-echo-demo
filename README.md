# kubernetes-workshop-echo-demo

deploy a demo application with built in config and with port 8081

## How to deploy

### Automatically with sed

Execute

```bash
sed -i "s/REPLACEME/$(whoami)/g" kustomization.yml ingress.yml
```

### Manually

Search for all occurences of `echo-demo.s.pub-dev.tech`. Replace the `echo-demo` with a unique name.

Do not change the `s.pub-dev.tech` part.

Change the namespace in the `kustomization.yaml` file to your namespace and run:

```bash

kubectl apply -k .

```

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
