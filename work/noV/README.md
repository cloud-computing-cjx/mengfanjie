## 作业
把我们的 httpserver 服务以 Istio Ingress Gateway 的形式发布出来。以下是你需要考虑的几点：
- 如何实现安全保证；
- 七层路由规则；
- 考虑 open tracing 的接入。

### 实现安全保证
#### deploy httpserver
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
#### check ingress ip
```
k get svc -nistio-system
istio-ingressgateway   LoadBalancer   10.97.5.233      <pending>     15021:31467/TCP,80:32392/TCP,443:32170/TCP,31400:30120/TCP,15443:30847/TCP   8d
```
#### access the httpserver via ingress
```
curl --resolve httpsserver.cncamp.io:443:10.97.5.233 https://httpsserver.cncamp.io/healthz -v -k

* Added httpsserver.cncamp.io:443:10.97.5.233 to DNS cache
* Hostname httpsserver.cncamp.io was found in DNS cache
*   Trying 10.97.5.233:443...
* TCP_NODELAY set
* Connected to httpsserver.cncamp.io (10.97.5.233) port 443 (#0)
* ALPN, offering h2
* ALPN, offering http/1.1
* successfully set certificate verify locations:
*   CAfile: /etc/ssl/certs/ca-certificates.crt
  CApath: /etc/ssl/certs
* TLSv1.3 (OUT), TLS handshake, Client hello (1):
* TLSv1.3 (IN), TLS handshake, Server hello (2):
* TLSv1.3 (IN), TLS handshake, Encrypted Extensions (8):
* TLSv1.3 (IN), TLS handshake, Certificate (11):
* TLSv1.3 (IN), TLS handshake, CERT verify (15):
* TLSv1.3 (IN), TLS handshake, Finished (20):
* TLSv1.3 (OUT), TLS change cipher, Change cipher spec (1):
* TLSv1.3 (OUT), TLS handshake, Finished (20):
* SSL connection using TLSv1.3 / TLS_AES_256_GCM_SHA384
* ALPN, server accepted to use h2
* Server certificate:
*  subject: O=cncamp Inc.; CN=*.cncamp.io
*  start date: Dec 26 15:18:36 2021 GMT
*  expire date: Dec 26 15:18:36 2022 GMT
*  issuer: O=cncamp Inc.; CN=*.cncamp.io
*  SSL certificate verify result: self signed certificate (18), continuing anyway.
* Using HTTP2, server supports multi-use
* Connection state changed (HTTP/2 confirmed)
* Copying HTTP/2 data in stream buffer to connection buffer after upgrade: len=0
* Using Stream ID: 1 (easy handle 0x556ba8687e30)
> GET /healthz HTTP/2
> Host: httpsserver.cncamp.io
> user-agent: curl/7.68.0
> accept: */*
> 
* TLSv1.3 (IN), TLS handshake, Newsession Ticket (4):
* TLSv1.3 (IN), TLS handshake, Newsession Ticket (4):
* old SSL session ID is stale, removing
* Connection state changed (MAX_CONCURRENT_STREAMS == 2147483647)!
< HTTP/2 200 
< date: Sun, 26 Dec 2021 15:39:41 GMT
< content-length: 3
< content-type: text/plain; charset=utf-8
< x-envoy-upstream-service-time: 24
< server: istio-envoy
< 
ok
* Connection #0 to host httpsserver.cncamp.io left intact
```

```
curl --resolve httpsserver.cncamp.io:443:10.97.5.233 https://httpsserver.cncamp.io/hello -v -k

* Added httpsserver.cncamp.io:443:10.97.5.233 to DNS cache
* Hostname httpsserver.cncamp.io was found in DNS cache
*   Trying 10.97.5.233:443...
* TCP_NODELAY set
* Connected to httpsserver.cncamp.io (10.97.5.233) port 443 (#0)
* ALPN, offering h2
* ALPN, offering http/1.1
* successfully set certificate verify locations:
*   CAfile: /etc/ssl/certs/ca-certificates.crt
  CApath: /etc/ssl/certs
* TLSv1.3 (OUT), TLS handshake, Client hello (1):
* TLSv1.3 (IN), TLS handshake, Server hello (2):
* TLSv1.3 (IN), TLS handshake, Encrypted Extensions (8):
* TLSv1.3 (IN), TLS handshake, Certificate (11):
* TLSv1.3 (IN), TLS handshake, CERT verify (15):
* TLSv1.3 (IN), TLS handshake, Finished (20):
* TLSv1.3 (OUT), TLS change cipher, Change cipher spec (1):
* TLSv1.3 (OUT), TLS handshake, Finished (20):
* SSL connection using TLSv1.3 / TLS_AES_256_GCM_SHA384
* ALPN, server accepted to use h2
* Server certificate:
*  subject: O=cncamp Inc.; CN=*.cncamp.io
*  start date: Dec 26 15:18:36 2021 GMT
*  expire date: Dec 26 15:18:36 2022 GMT
*  issuer: O=cncamp Inc.; CN=*.cncamp.io
*  SSL certificate verify result: self signed certificate (18), continuing anyway.
* Using HTTP2, server supports multi-use
* Connection state changed (HTTP/2 confirmed)
* Copying HTTP/2 data in stream buffer to connection buffer after upgrade: len=0
* Using Stream ID: 1 (easy handle 0x55c4b8c08e30)
> GET /hello HTTP/2
> Host: httpsserver.cncamp.io
> user-agent: curl/7.68.0
> accept: */*
> 
* TLSv1.3 (IN), TLS handshake, Newsession Ticket (4):
* TLSv1.3 (IN), TLS handshake, Newsession Ticket (4):
* old SSL session ID is stale, removing
* Connection state changed (MAX_CONCURRENT_STREAMS == 2147483647)!
< HTTP/2 200 
< date: Sun, 26 Dec 2021 15:50:43 GMT
< content-length: 657
< content-type: text/plain; charset=utf-8
< x-envoy-upstream-service-time: 140
< server: istio-envoy
< 
hello [stranger]
===================Details of the http request header:============
Accept=[*/*]
X-Request-Id=[fb8877c2-a883-9329-a21f-a807fdf1e4ed]
X-Forwarded-Client-Cert=[By=spiffe://cluster.local/ns/httpsserver/sa/default;Hash=2f18708b4b86fe81aa2728013357fcf98067bd848b64140c1ea8f2462bfe813b;Subject="";URI=spiffe://cluster.local/ns/istio-system/sa/istio-ingressgateway-service-account]
X-B3-Spanid=[3941b262db94b5ea]
X-B3-Sampled=[1]
User-Agent=[curl/7.68.0]
X-Forwarded-For=[192.168.3.150]
X-Forwarded-Proto=[https]
X-Envoy-Attempt-Count=[1]
X-Envoy-Internal=[true]
X-B3-Traceid=[0aa71791dacf6c63e347bc24b07d07d8]
X-B3-Parentspanid=[e347bc24b07d07d8]
* Connection #0 to host httpsserver.cncamp.io left intact
```
### 七层路由规则
```
curl --resolve httpsserver.cncamp.io:443:10.97.5.233 https://httpsserver.cncamp.io/httpsserver/hello -v -k
```
### open tracing 的接入
#### install
```
kubectl apply -f jaeger.yaml
kubectl edit configmap istio -n istio-system
set tracing.sampling=100
```
```
# Please edit the object below. Lines beginning with a '#' will be ignored,
# and an empty file will abort the edit. If an error occurs while saving this file will be
# reopened with the relevant failures.
#
apiVersion: v1
data:
  mesh: |-
    accessLogFile: /dev/stdout
    defaultConfig:
      discoveryAddress: istiod.istio-system.svc:15012
      proxyMetadata: {}
      tracing:
        sampling: 100
        zipkin:
          address: zipkin.istio-system:9411
    enablePrometheusMerge: true
    rootNamespace: istio-system
    trustDomain: cluster.local
  meshNetworks: 'networks: {}'
kind: ConfigMap
metadata:
  annotations:
    kubectl.kubernetes.io/last-applied-configuration: |
      {"apiVersion":"v1","data":{"mesh":"accessLogFile: /dev/stdout\ndefaultConfig:\n  discoveryAddress: istiod.istio-system.svc:15012\n  proxyMetadata: {}\n  tracing:\n    zipkin:\n      address: zipkin.istio-system:9411\nenablePrometheusMerge: true\nrootNamespace: istio-system\ntrustDomain: cluster.local","meshNetworks":"networks: {}"},"kind":"ConfigMap","metadata":{"annotations":{},"labels":{"install.operator.istio.io/owning-resource":"unknown","install.operator.istio.io/owning-resource-namespace":"istio-system","istio.io/rev":"default","operator.istio.io/component":"Pilot","operator.istio.io/managed":"Reconcile","operator.istio.io/version":"1.12.1","release":"istio"},"name":"istio","namespace":"istio-system"}}
  creationTimestamp: "2021-12-18T14:20:17Z"
  labels:
    install.operator.istio.io/owning-resource: unknown
    install.operator.istio.io/owning-resource-namespace: istio-system
    istio.io/rev: default
    operator.istio.io/component: Pilot
    operator.istio.io/managed: Reconcile
    operator.istio.io/version: 1.12.1
    release: istio
  name: istio
  namespace: istio-system
  resourceVersion: "216871"
  uid: 0d34255e-a416-44e2-9911-9994b4b383f0
```
#### deploy istiodemo
```
kubectl create ns istiodemo
kubectl label ns istiodemo istio-injection=enabled
kubectl create -f service0.yaml -n istiodemo
kubectl create -f service1.yaml -n istiodemo
kubectl create -f service2.yaml -n istiodemo
kubectl create -f istio-specs-tracing.yaml -n istiodemo
```
#### check ingress ip
```
k get svc -nistio-system
istio-ingressgateway   LoadBalancer   10.97.5.233      <pending>     15021:31467/TCP,80:32392/TCP,443:32170/TCP,31400:30120/TCP,15443:30847/TCP   8d
```
#### access the istiodemo via ingress for 100 times
```
root@kubemaster:~/go/src/github.com/cncamp/mengfanjie/work/noV# curl 10.97.5.233/service0
===================Details of the http request header:============
HTTP/1.1 200 OK
Transfer-Encoding: chunked
Content-Type: text/plain; charset=utf-8
Date: Sun, 26 Dec 2021 16:01:40 GMT
Server: envoy
X-Envoy-Upstream-Service-Time: 126

38a
===================Details of the http request header:============
HTTP/1.1 200 OK
Content-Length: 671
Content-Type: text/plain; charset=utf-8
Date: Sun, 26 Dec 2021 16:01:39 GMT
Server: envoy
X-Envoy-Upstream-Service-Time: 65

===================Details of the http request header:============
User-Agent=[Go-http-client/1.1,Go-http-client/1.1,curl/7.68.0]
Accept=[*/*]
X-Forwarded-For=[192.168.3.150]
X-Request-Id=[f7a9a9da-a184-9f01-b5b3-4fa40c8399cf]
X-B3-Traceid=[aa9b0e69c6a9c6972b7c4380e8e21c15]
Accept-Encoding=[gzip,gzip]
X-Forwarded-Proto=[http]
X-Forwarded-Client-Cert=[By=spiffe://cluster.local/ns/istiodemo/sa/default;Hash=608c4f25c23b4263e690b00a90ff153105b68277cc3ef42e39551789ab676210;Subject="";URI=spiffe://cluster.local/ns/istiodemo/sa/default]
X-B3-Sampled=[1]
X-Envoy-Attempt-Count=[1]
X-Envoy-Internal=[true]
X-B3-Spanid=[538c4b5b3cb1f576]
X-B3-Parentspanid=[5cc1b9ef18e21802]

0
```

```
root@kubemaster:~/go/src/github.com/cncamp/mengfanjie/work/noV# k get po -n istiodemo
NAME                        READY   STATUS    RESTARTS   AGE
service0-865968847b-898nk   2/2     Running   0          5m7s
service1-7c9df986b9-pjtwj   2/2     Running   0          5m3s
service2-6d6cc75fd-gbjgc    2/2     Running   0          5m
root@kubemaster:~/go/src/github.com/cncamp/mengfanjie/work/noV# k logs -f service2-6d6cc75fd-gbjgc -n istiodemo
ERROR: logging before flag.Parse: I1226 15:58:54.338103       1 main.go:21] Starting service2
ERROR: logging before flag.Parse: I1226 15:58:54.338210       1 main.go:38] Server Started
ERROR: logging before flag.Parse: I1226 16:01:39.982395       1 main.go:67] entering v2 root handler
ERROR: logging before flag.Parse: I1226 16:01:39.999767       1 main.go:75] Respond in 17 ms

```
#### check tracing dashboard
```
root@kubemaster:~/go/src/github.com/cncamp/mengfanjie/work/noV# istioctl dashboard jaeger
http://localhost:16686
```