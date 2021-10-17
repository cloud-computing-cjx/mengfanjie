# 构建本地镜像。
## 编写 Dockerfile 将练习 2.2 编写的 httpserver 容器化（请思考有哪些最佳实践可以引入到 Dockerfile 中来）。
```
➜  noII git:(main) ✗ docker build -t my-go-http .
```
## 将镜像推送至 Docker 官方镜像仓库。
```
➜  noII git:(main) ✗ docker tag my-go-http:latest jinyumantang/httpserver:v1.20211017.18.43
➜  noII git:(main) ✗ docker login
➜  noII git:(main) ✗ docker push jinyumantang/httpserver:v1.20211017.18.43
```
https://hub.docker.com/repository/docker/jinyumantang/httpserver
## 通过 Docker 命令本地启动 httpserver。
```
➜  noII git:(main) ✗ docker run -it -p 8000:8000 --rm --name httpserver my-go-http
```
## 通过 nsenter 进入容器查看 IP 配置。
```
➜  ~ docker ps
CONTAINER ID   IMAGE        COMMAND                  CREATED          STATUS          PORTS                                       NAMES
462dcc796b60   my-go-http   "go run /go/src/http…"   43 seconds ago   Up 42 seconds   0.0.0.0:8000->8000/tcp, :::8000->8000/tcp   httpserver
➜  ~ docker inspect 462dcc796b60|grep -i pid
            "Pid": 25185,
            "PidMode": "",
            "PidsLimit": null,
➜  ~ sudo nsenter -t 25185 -n ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
11: eth0@if12: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default 
    link/ether 02:42:ac:11:00:02 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 172.17.0.2/16 brd 172.17.255.255 scope global eth0
       valid_lft forever preferred_lft forever
```