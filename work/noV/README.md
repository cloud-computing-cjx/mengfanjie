## 作业
把我们的 httpserver 服务以 Istio Ingress Gateway 的形式发布出来。以下是你需要考虑的几点：
- 如何实现安全保证；
- 七层路由规则；
- 考虑 open tracing 的接入。

## deploy httpserver
```
kubectl create ns httpsserver
kubectl label ns httpsserver istio-injection=enabled
kubectl create -f httpserver.yaml -n httpsserver
```
```
openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -subj '/O=cncamp Inc./CN=*.cncamp.io' -keyout cncamp.io.key -out cncamp.io.crt
kubectl create -n istio-system secret tls cncamp-credential --key=cncamp.io.key --cert=cncamp.io.crt
kubectl create -f istio-specs.yaml -n httpsserver
```

### check ingress ip
```
k get svc -nistio-system
istio-ingressgateway   LoadBalancer   10.97.5.233      <pending>     15021:31467/TCP,80:32392/TCP,443:32170/TCP,31400:30120/TCP,15443:30847/TCP   8d
```
### access the httpserver via ingress
```
curl --resolve httpsserver.cncamp.io:443:10.105.227.152 https://httpsserver.cncamp.io/healthz -v -k

curl --resolve httpsserver.cncamp.io:443:10.97.5.233 https://httpsserver.cncamp.io/healthz -v -k
```
