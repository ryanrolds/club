# K8s files

Docker Build

```
docker tag 4ea6105f31a3 docker.pedanticorderliness.com/club:latest
docker push docker.pedanticorderliness.com/club:latest
```

K8s

```
export ENV=test
export TAG_NAME=latest
envsubst < deployment.yaml | kubectl apply -f -
envsubst < service.yaml | kubectl apply -f -
kubectl apply -f ingress-test.yaml
```

