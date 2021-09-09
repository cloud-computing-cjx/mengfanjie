# 作业
## 搭建环境测试，通过static pod的方式启动应用
要获知版本信息，请输入 kubectl version.
### 创建静态 Pod
#### 文件系统上的静态 Pod 声明文件
##### 选择一个要运行静态 Pod 的节点。
##### 选择一个目录，比如在 /etc/kubelet.d 目录来保存 web 服务 Pod 的定义文件， /etc/kubelet.d/static-web.yaml：
```
# 在 kubelet 运行的节点上执行以下命令
mkdir /etc/kubelet.d/
cat <<EOF >/etc/kubelet.d/static-web.yaml
apiVersion: v1
kind: Pod
metadata:
  name: static-web
  labels:
    role: myrole
spec:
  containers:
    - name: web
      image: nginx
      ports:
        - name: web
          containerPort: 80
          protocol: TCP
EOF
```
##### 配置这个节点上的 kubelet，使用这个参数执行 --pod-manifest-path=/etc/kubelet.d/。 在 Fedora 上编辑 /etc/kubernetes/kubelet 以包含下行：
```
KUBELET_ARGS="--cluster-dns=10.254.0.10 --cluster-domain=kube.local --pod-manifest-path=/etc/kubelet.d/"
```
##### 重启 kubelet。Fedora 上使用下面的命令：
```
# 在 kubelet 运行的节点上执行以下命令
systemctl restart kubelet
```
#### Web 网上的静态 Pod 声明文件
