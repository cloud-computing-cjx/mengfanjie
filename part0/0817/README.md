# 编写 yaml
```
[root@bigdata-pro04 ~]# mkdir pods
[root@bigdata-pro04 ~]# cd pods/
```
redisserver.yaml
```
[root@bigdata-pro04 pods]# vim redisserver.yaml

apiVersion: apps/v1
kind: Deployment
metadata:
    name: redisserver
    namespace: default
spec:
    replicas: 2
    selector:
        matchLabels:
            run: redisserver
    template:
        metadata:
            labels:
                run: redisserver
        spec:
            containers:
            - image: redis:latest
              name: redisserver
```
nginx.yaml
```
[root@bigdata-pro04 pods]# vim nginx.yaml

apiVersion: apps/v1
kind: Deployment
metadata:
    name: nginx
    namespace: default
spec:
    replicas: 2
    selector:
        matchLabels:
            run: nginx
    template:
        metadata:
            labels:
                run: nginx
        spec:
            containers:
            - image: nginx:latest
              name: nginx
```
# 运行 pod
redisserver.yaml 
```
[root@bigdata-pro04 pods]# kubectl create -f redisserver.yaml 
deployment.apps/redisserver created
[root@bigdata-pro04 pods]# kubectl get pod -owide
NAME                           READY   STATUS              RESTARTS   AGE   IP       NODE            NOMINATED NODE   READINESS GATES
redisserver-7dd558df44-rn4gt   0/1     ContainerCreating   0          5s    <none>   bigdata-pro03   <none>           <none>
redisserver-7dd558df44-sg7b4   0/1     ContainerCreating   0          5s    <none>   bigdata-pro03   <none>           <none>
```
nginx.yaml
```
[root@bigdata-pro04 pods]# kubectl create -f nginx.yaml 
deployment.apps/nginx created
[root@bigdata-pro04 pods]# kubectl get pod -owide
NAME                           READY   STATUS              RESTARTS   AGE     IP              NODE            NOMINATED NODE   READINESS GATES
nginx-5fc7bcd4d9-c9cn6         0/1     ContainerCreating   0          10s     <none>          bigdata-pro03   <none>           <none>
nginx-5fc7bcd4d9-ldrfm         0/1     ContainerCreating   0          10s     <none>          bigdata-pro03   <none>           <none>
redisserver-7dd558df44-rn4gt   1/1     Running             0          9m33s   10.100.144.66   bigdata-pro03   <none>           <none>
redisserver-7dd558df44-sg7b4   1/1     Running             0          9m33s   10.100.144.65   bigdata-pro03   <none>           <none>
```
访问 nginx
```
[root@bigdata-pro04 pods]# curl 10.100.144.67
......
```
## create 和 apply(常用) 的区别
查看日志
```
kubectl create -f nginx.yaml -v 9
kubectl apply -f nginx.yaml -v 9
```
- create 最终执行的是: curl -k -v -XPOST ...
- apply 最终执行的是: curl -k -v -XPATCH ...

PATCH 方法只发送增量部分，POST 发送全量
# 更新 pod
```
[root@bigdata-pro04 pods]# kubectl scale deploy nginx --replicas=3
deployment.apps/nginx scaled
[root@bigdata-pro04 pods]# kubectl get pod -owide
NAME                           READY   STATUS              RESTARTS   AGE     IP              NODE            NOMINATED NODE   READINESS GATES
nginx-5fc7bcd4d9-c9cn6         1/1     Running             0          4m30s   10.100.144.67   bigdata-pro03   <none>           <none>
nginx-5fc7bcd4d9-ldrfm         1/1     Running             0          4m30s   10.100.144.68   bigdata-pro03   <none>           <none>
nginx-5fc7bcd4d9-xv57k         0/1     ContainerCreating   0          6s      <none>          bigdata-pro03   <none>           <none>
redisserver-7dd558df44-rn4gt   1/1     Running             0          13m     10.100.144.66   bigdata-pro03   <none>           <none>
redisserver-7dd558df44-sg7b4   1/1     Running             0          13m     10.100.144.65   bigdata-pro03   <none>           <none>
```
# 删除 pod
```
[root@bigdata-pro04 pods]# kubectl delete -f nginx.yaml 
deployment.apps "nginx" deleted
[root@bigdata-pro04 pods]# kubectl get pod -owide
NAME                           READY   STATUS        RESTARTS   AGE     IP              NODE            NOMINATED NODE   READINESS GATES
nginx-5fc7bcd4d9-c9cn6         0/1     Terminating   0          13m     10.100.144.67   bigdata-pro03   <none>           <none>
nginx-5fc7bcd4d9-ldrfm         0/1     Terminating   0          13m     10.100.144.68   bigdata-pro03   <none>           <none>
nginx-5fc7bcd4d9-xv57k         0/1     Terminating   0          9m13s   10.100.144.69   bigdata-pro03   <none>           <none>
redisserver-7dd558df44-rn4gt   1/1     Running       0          23m     10.100.144.66   bigdata-pro03   <none>           <none>
redisserver-7dd558df44-sg7b4   1/1     Running       0          23m     10.100.144.65   bigdata-pro03   <none>           <none>
[root@bigdata-pro04 pods]# kubectl get pod -owide
NAME                           READY   STATUS    RESTARTS   AGE   IP              NODE            NOMINATED NODE   READINESS GATES
redisserver-7dd558df44-rn4gt   1/1     Running   0          23m   10.100.144.66   bigdata-pro03   <none>           <none>
redisserver-7dd558df44-sg7b4   1/1     Running   0          23m   10.100.144.65   bigdata-pro03   <none>           <none>
```
# Server
创建 pod
```
[root@bigdata-pro04 pods]# vim run-my-nginx.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-nginx
spec:
  selector:
    matchLabels:
      run: my-nginx
  replicas: 2
  template:
    metadata:
      labels:
        run: my-nginx
    spec:
      containers:
      - name: my-nginx
        image: nginx
        ports:
        - containerPort: 80

```
启动 pod
```
[root@bigdata-pro04 pods]# kubectl apply -f ./run-my-nginx.yaml
deployment.apps/my-nginx created
[root@bigdata-pro04 pods]# kubectl get pods -l run=my-nginx -o wide
NAME                        READY   STATUS              RESTARTS   AGE   IP       NODE            NOMINATED NODE   READINESS GATES
my-nginx-5b56ccd65f-cf7pt   0/1     ContainerCreating   0          7s    <none>   bigdata-pro03   <none>           <none>
my-nginx-5b56ccd65f-sjbl9   0/1     ContainerCreating   0          7s    <none>   bigdata-pro03   <none>           <none>
```
检查 Pod 的 IP 地址：
```
[root@bigdata-pro04 pods]# kubectl get pods -l run=my-nginx -o yaml | grep podIP
      cni.projectcalico.org/podIP: 10.100.144.70/32
      cni.projectcalico.org/podIPs: 10.100.144.70/32
    podIP: 10.100.144.70
    podIPs:
      cni.projectcalico.org/podIP: 10.100.144.71/32
      cni.projectcalico.org/podIPs: 10.100.144.71/32
    podIP: 10.100.144.71
    podIPs:
```
## 创建 Service
```
[root@bigdata-pro04 pods]# kubectl expose deployment/my-nginx
service/my-nginx exposed
```
这等价于使用 kubectl create -f 命令创建，对应如下的 yaml 文件：
```
[root@bigdata-pro04 pods]# mkdir -p service/networking
[root@bigdata-pro04 pods]# vim service/networking/nginx-svc.yaml

apiVersion: v1
kind: Service
metadata:
  name: my-nginx
  labels:
    run: my-nginx
spec:
  ports:
  - port: 80
    protocol: TCP
  selector:
    run: my-nginx

```
## 查看 Service
```
[root@bigdata-pro04 pods]# kubectl get svc
NAME         TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)   AGE
kubernetes   ClusterIP   10.96.0.1       <none>        443/TCP   2d5h
my-nginx     ClusterIP   10.96.173.184   <none>        80/TCP    3m32s
[root@bigdata-pro04 pods]# kubectl get svc my-nginx
NAME       TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)   AGE
my-nginx   ClusterIP   10.96.173.184   <none>        80/TCP    3m25s

[root@bigdata-pro04 pods]# kubectl describe svc my-nginx
Name:              my-nginx
Namespace:         default
Labels:            <none>
Annotations:       <none>
Selector:          run=my-nginx
Type:              ClusterIP
IP Family Policy:  SingleStack
IP Families:       IPv4
IP:                10.96.173.184
IPs:               10.96.173.184
Port:              <unset>  80/TCP
TargetPort:        80/TCP
Endpoints:         10.100.144.70:80,10.100.144.71:80
Session Affinity:  None
Events:            <none>

[root@bigdata-pro04 pods]# kubectl get ep my-nginx
NAME       ENDPOINTS                           AGE
my-nginx   10.100.144.70:80,10.100.144.71:80   5m20s

```
测试
```
[root@bigdata-pro04 pods]# curl 10.96.173.184
...
```
## Service 详情
```
[root@bigdata-pro04 pods]# kubectl get svc -oyaml
...
```
# Kubernetes对象的通用设计
```
[root@bigdata-pro04 pods]# kubectl get pod my-nginx-5b56ccd65f-cf7pt -oyaml
......

```
## labels
主要做过滤查询
```
[root@bigdata-pro04 pods]# kubectl get pod --show-labels
NAME                           READY   STATUS    RESTARTS   AGE   LABELS
my-nginx-5b56ccd65f-cf7pt      1/1     Running   0          20m   pod-template-hash=5b56ccd65f,run=my-nginx
my-nginx-5b56ccd65f-sjbl9      1/1     Running   0          20m   pod-template-hash=5b56ccd65f,run=my-nginx
redisserver-7dd558df44-rn4gt   1/1     Running   0          44m   pod-template-hash=7dd558df44,run=redisserver
redisserver-7dd558df44-sg7b4   1/1     Running   0          44m   pod-template-hash=7dd558df44,run=redisserver

[root@bigdata-pro04 pods]# kubectl get pod -l run=my-nginx
NAME                        READY   STATUS    RESTARTS   AGE
my-nginx-5b56ccd65f-cf7pt   1/1     Running   0          21m
my-nginx-5b56ccd65f-sjbl9   1/1     Running   0          21m
```
label 可以 配合 service 的 spec.selector 完成一些业务
## annotation
主要作为注解
### 添加 annotate
```
[root@bigdata-pro04 pods]# kubectl annotate pod my-nginx-5b56ccd65f-sjbl9 a=b
pod/my-nginx-5b56ccd65f-sjbl9 annotated

[root@bigdata-pro04 ~]# kubectl get pod my-nginx-5b56ccd65f-sjbl9 -oyaml
```
## Finalizers
资源的锁（加锁后不会被删除）
## ResourceVersion
版本控制（乐观锁，MVCC）
## SelftLink

# 部署（Deployment）
```
[root@bigdata-pro04 pods]# kubectl get pod
NAME                           READY   STATUS    RESTARTS   AGE
my-nginx-5b56ccd65f-cf7pt      1/1     Running   0          59m
my-nginx-5b56ccd65f-sjbl9      1/1     Running   0          59m
redisserver-7dd558df44-rn4gt   1/1     Running   0          84m
redisserver-7dd558df44-sg7b4   1/1     Running   0          84m
[root@bigdata-pro04 pods]# kubectl get deploy
NAME          READY   UP-TO-DATE   AVAILABLE   AGE
my-nginx      2/2     2            2           59m
redisserver   2/2     2            2           84m
```
## 查看详细信息
```
[root@bigdata-pro04 pods]# kubectl get deploy my-nginx -oyaml
...
```
默认的升级策略
```
spec:
    strategy:
        rollingUpdate:
            maxSurge: 25%
            maxUnavailable: 25%
        type: RollingUpdate
```
描述
```
[root@bigdata-pro04 pods]# kubectl describe deploy my-nginx
...
```
## 升级
```
[root@bigdata-pro04 pods]# kubectl get deploy
NAME          READY   UP-TO-DATE   AVAILABLE   AGE
my-nginx      2/2     2            2           74m
redisserver   2/2     2            2           99m

[root@bigdata-pro04 pods]# kubectl edit deploy my-nginx

[root@bigdata-pro04 pods]# kubectl get rs
更新中
[root@bigdata-pro04 pods]# kubectl get rs -w
更新过程
[root@bigdata-pro04 pods]# kubectl get pod
滚动升级完成
```
# Try it
通过类似 docker run 的命令在 Kubernetes 运行容器
```
[root@bigdata-pro04 pods]# kubectl run --image=nginx:alpine nginx-app --port=80
pod/nginx-app created

[root@bigdata-pro04 pods]# kubectl get deployment
NAME          READY   UP-TO-DATE   AVAILABLE   AGE
my-nginx      2/2     2            2           92m
redisserver   2/2     2            2           117m

[root@bigdata-pro04 pods]# kubectl get pod
NAME                           READY   STATUS    RESTARTS   AGE
my-nginx-5b56ccd65f-cf7pt      1/1     Running   0          93m
my-nginx-5b56ccd65f-sjbl9      1/1     Running   0          93m
nginx-app                      1/1     Running   0          17s
redisserver-7dd558df44-rn4gt   1/1     Running   0          117m
redisserver-7dd558df44-sg7b4   1/1     Running   0          117m
```