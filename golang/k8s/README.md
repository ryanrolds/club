# K8s files

Docker Build

```
docker build .
<get final tag (4ea6105f31a3) from output >
docker tag 4ea6105f31a3 docker.pedanticorderliness.com/club-backend:4ea6105f31a3
docker push docker.pedanticorderliness.com/club-backend:4ea6105f31a3
```

K8s

```
export ENV=prod
export TAG_NAME=4ea6105f31a3
envsubst < deployment.yaml | kubectl apply -f -
envsubst < service.yaml | kubectl apply -f -
kubectl apply -f ingress-prod.yaml
```

### Misc

Updating Docker "regcreds".

Set `htpassword` in `docker-registry-secret` secret. The value is a base64 encoded bcrypted htpassword string.

```
htpasswd -Bnb <admin> <password >| base64 -w 0
<base64 bcrypted htpassword string>
kubectl edit secret kubectl edit secret docker-registry-secret
<restart Docker registry>
```

Update `regcred` containing condentials K8s needs to get Docker iamges from private repo.

```
kubectl create secret docker-registry regcred --docker-server=docker.pedanticorderliness.com --docker-username=<username> --docker-password=<password> --docker-email=<email>
```
