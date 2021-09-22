# k8s操作
## 节点信息
```
[root@bigdata-pro04 ~]# kubectl get nodes
NAME            STATUS   ROLES                  AGE    VERSION
bigdata-pro03   Ready    <none>                 47h    v1.21.4
bigdata-pro04   Ready    control-plane,master   2d3h   v1.21.4
[root@bigdata-pro04 ~]# kubectl get nodes -o wide
NAME            STATUS   ROLES                  AGE    VERSION   INTERNAL-IP     EXTERNAL-IP   OS-IMAGE                KERNEL-VERSION                CONTAINER-RUNTIME
bigdata-pro03   Ready    <none>                 47h    v1.21.4   192.168.3.153   <none>        CentOS Linux 7 (Core)   3.10.0-1160.36.2.el7.x86_64   docker://1.13.1
bigdata-pro04   Ready    control-plane,master   2d3h   v1.21.4   192.168.3.154   <none>        CentOS Linux 7 (Core)   3.10.0-1160.15.2.el7.x86_64   docker://1.13.1
```
## 正在运行的服务组件
```
[root@bigdata-pro04 ~]# kubectl get pod -n kube-system
NAME                                    READY   STATUS    RESTARTS   AGE
coredns-7d75679df-khphl                 1/1     Running   0          2d3h
coredns-7d75679df-xdvrx                 1/1     Running   0          2d3h
etcd-bigdata-pro04                      1/1     Running   0          2d3h
kube-apiserver-bigdata-pro04            1/1     Running   0          2d3h
kube-controller-manager-bigdata-pro04   1/1     Running   0          2d3h
kube-proxy-6bgn4                        1/1     Running   0          47h
kube-proxy-tt9zf                        1/1     Running   0          2d3h
kube-scheduler-bigdata-pro04            1/1     Running   0          2d3h
```
## etcd-bigdata-pro04 的详细信息
```
[root@bigdata-pro04 ~]# kubectl get pod etcd-bigdata-pro04 -n kube-system -oyaml
apiVersion: v1
kind: Pod
metadata:
  annotations:
    kubeadm.kubernetes.io/etcd.advertise-client-urls: https://192.168.3.154:2379
    kubernetes.io/config.hash: ea8b8fc7ee185dd2b42e3376d1c99646
    kubernetes.io/config.mirror: ea8b8fc7ee185dd2b42e3376d1c99646
    kubernetes.io/config.seen: "2021-08-20T19:14:09.312448987+08:00"
    kubernetes.io/config.source: file
  creationTimestamp: "2021-08-20T11:14:25Z"
  labels:
    component: etcd
    tier: control-plane
  name: etcd-bigdata-pro04
  namespace: kube-system
  ownerReferences:
  - apiVersion: v1
    controller: true
    kind: Node
    name: bigdata-pro04
    uid: 69e4f59a-7de2-457e-b835-203294fc4526
  resourceVersion: "365"
  uid: 8faa56db-f0d6-43a0-8431-1935b2923f92
spec:
  containers:
  - command:
    - etcd
    - --advertise-client-urls=https://192.168.3.154:2379
    - --cert-file=/etc/kubernetes/pki/etcd/server.crt
    - --client-cert-auth=true
    - --data-dir=/var/lib/etcd
    - --initial-advertise-peer-urls=https://192.168.3.154:2380
    - --initial-cluster=bigdata-pro04=https://192.168.3.154:2380
    - --key-file=/etc/kubernetes/pki/etcd/server.key
    - --listen-client-urls=https://127.0.0.1:2379,https://192.168.3.154:2379
    - --listen-metrics-urls=http://127.0.0.1:2381
    - --listen-peer-urls=https://192.168.3.154:2380
    - --name=bigdata-pro04
    - --peer-cert-file=/etc/kubernetes/pki/etcd/peer.crt
    - --peer-client-cert-auth=true
    - --peer-key-file=/etc/kubernetes/pki/etcd/peer.key
    - --peer-trusted-ca-file=/etc/kubernetes/pki/etcd/ca.crt
    - --snapshot-count=10000
    - --trusted-ca-file=/etc/kubernetes/pki/etcd/ca.crt
    image: registry.aliyuncs.com/k8sxio/etcd:3.4.13-0
    imagePullPolicy: IfNotPresent
    livenessProbe:
      failureThreshold: 8
      httpGet:
        host: 127.0.0.1
        path: /health
        port: 2381
        scheme: HTTP
      initialDelaySeconds: 10
      periodSeconds: 10
      successThreshold: 1
      timeoutSeconds: 15
    name: etcd
    resources:
      requests:
        cpu: 100m
        memory: 100Mi
    startupProbe:
      failureThreshold: 24
      httpGet:
        host: 127.0.0.1
        path: /health
        port: 2381
        scheme: HTTP
      initialDelaySeconds: 10
      periodSeconds: 10
      successThreshold: 1
      timeoutSeconds: 15
    terminationMessagePath: /dev/termination-log
    terminationMessagePolicy: File
    volumeMounts:
    - mountPath: /var/lib/etcd
      name: etcd-data
    - mountPath: /etc/kubernetes/pki/etcd
      name: etcd-certs
  dnsPolicy: ClusterFirst
  enableServiceLinks: true
  hostNetwork: true
  nodeName: bigdata-pro04
  preemptionPolicy: PreemptLowerPriority
  priority: 2000001000
  priorityClassName: system-node-critical
  restartPolicy: Always
  schedulerName: default-scheduler
  securityContext: {}
  terminationGracePeriodSeconds: 30
  tolerations:
  - effect: NoExecute
    operator: Exists
  volumes:
  - hostPath:
      path: /etc/kubernetes/pki/etcd
      type: DirectoryOrCreate
    name: etcd-certs
  - hostPath:
      path: /var/lib/etcd
      type: DirectoryOrCreate
    name: etcd-data
status:
  conditions:
  - lastProbeTime: null
    lastTransitionTime: "2021-08-20T11:14:15Z"
    status: "True"
    type: Initialized
  - lastProbeTime: null
    lastTransitionTime: "2021-08-20T11:14:37Z"
    status: "True"
    type: Ready
  - lastProbeTime: null
    lastTransitionTime: "2021-08-20T11:14:37Z"
    status: "True"
    type: ContainersReady
  - lastProbeTime: null
    lastTransitionTime: "2021-08-20T11:14:15Z"
    status: "True"
    type: PodScheduled
  containerStatuses:
  - containerID: docker://fdc8eb773b27b77d4a12ff8ef76d90195c7269a890dba6708c8da1aa1cee8538
    image: registry.aliyuncs.com/k8sxio/etcd:3.4.13-0
    imageID: docker-pullable://registry.aliyuncs.com/k8sxio/etcd@sha256:bd4d2c9a19be8a492bc79df53eee199fd04b415e9993eb69f7718052602a147a
    lastState: {}
    name: etcd
    ready: true
    restartCount: 0
    started: true
    state:
      running:
        startedAt: "2021-08-20T11:14:17Z"
  hostIP: 192.168.3.154
  phase: Running
  podIP: 192.168.3.154
  podIPs:
  - ip: 192.168.3.154
  qosClass: Burstable
  startTime: "2021-08-20T11:14:15Z"
```
### etcd
#### 直接访问etcd的数据
etcd内部没有启动shell，所以无法进入交互 
```
[root@bigdata-pro04 ~]# ps -ef|grep etcd
root      9791  9764 13 8月20 ?       06:56:56 kube-apiserver --advertise-address=192.168.3.154 --allow-privileged=true --authorization-mode=Node,RBAC --client-ca-file=/etc/kubernetes/pki/ca.crt --enable-admission-plugins=NodeRestriction --enable-bootstrap-token-auth=true --etcd-cafile=/etc/kubernetes/pki/etcd/ca.crt --etcd-certfile=/etc/kubernetes/pki/apiserver-etcd-client.crt --etcd-keyfile=/etc/kubernetes/pki/apiserver-etcd-client.key --etcd-servers=https://127.0.0.1:2379 --insecure-port=0 --kubelet-client-certificate=/etc/kubernetes/pki/apiserver-kubelet-client.crt --kubelet-client-key=/etc/kubernetes/pki/apiserver-kubelet-client.key --kubelet-preferred-address-types=InternalIP,ExternalIP,Hostname --proxy-client-cert-file=/etc/kubernetes/pki/front-proxy-client.crt --proxy-client-key-file=/etc/kubernetes/pki/front-proxy-client.key --requestheader-allowed-names=front-proxy-client --requestheader-client-ca-file=/etc/kubernetes/pki/front-proxy-ca.crt --requestheader-extra-headers-prefix=X-Remote-Extra- --requestheader-group-headers=X-Remote-Group --requestheader-username-headers=X-Remote-User --secure-port=6443 --service-account-issuer=https://kubernetes.default.svc.cluster.local --service-account-key-file=/etc/kubernetes/pki/sa.pub --service-account-signing-key-file=/etc/kubernetes/pki/sa.key --service-cluster-ip-range=10.96.0.0/16 --tls-cert-file=/etc/kubernetes/pki/apiserver.crt --tls-private-key-file=/etc/kubernetes/pki/apiserver.key
root      9870  9801  4 8月20 ?       02:13:16 etcd --advertise-client-urls=https://192.168.3.154:2379 --cert-file=/etc/kubernetes/pki/etcd/server.crt --client-cert-auth=true --data-dir=/var/lib/etcd --initial-advertise-peer-urls=https://192.168.3.154:2380 --initial-cluster=bigdata-pro04=https://192.168.3.154:2380 --key-file=/etc/kubernetes/pki/etcd/server.key --listen-client-urls=https://127.0.0.1:2379,https://192.168.3.154:2379 --listen-metrics-urls=http://127.0.0.1:2381 --listen-peer-urls=https://192.168.3.154:2380 --name=bigdata-pro04 --peer-cert-file=/etc/kubernetes/pki/etcd/peer.crt --peer-client-cert-auth=true --peer-key-file=/etc/kubernetes/pki/etcd/peer.key --peer-trusted-ca-file=/etc/kubernetes/pki/etcd/ca.crt --snapshot-count=10000 --trusted-ca-file=/etc/kubernetes/pki/etcd/ca.crt
root     15142 10947  0 22:35 pts/1    00:00:00 grep --color=auto etcd
[root@bigdata-pro04 ~]# docker ps|grep etcd
fdc8eb773b27        0369cf4303ff                                                                                                                                  "etcd --advertise-..."   2 days ago          Up 2 days                               k8s_etcd_etcd-bigdata-pro04_kube-system_ea8b8fc7ee185dd2b42e3376d1c99646_0
08e68c40cf59        registry.aliyuncs.com/k8sxio/pause:3.4.1                                                                                                      "/pause"                 2 days ago          Up 2 days                               k8s_POD_etcd-bigdata-pro04_kube-system_ea8b8fc7ee185dd2b42e3376d1c99646_0
[root@bigdata-pro04 ~]# docker exec -it fdc8eb773b27 sh
sh-5.0#                                                           
```
##### 访问etcd的数据
```
设置API版本
sh-5.0# export ETCDCTL_API=3

sh-5.0# etcdctl --endpoints https://localhost:2379 --cert /etc/kubernetes/pki/etcd/server.crt --key /etc/kubernetes/pki/etcd/server.key --cacert /etc/kubernetes/pki/etcd/ca.crt get --keys-only --prefix /
/registry/apiextensions.k8s.io/customresourcedefinitions/bgpconfigurations.crd.projectcalico.org

/registry/apiextensions.k8s.io/customresourcedefinitions/bgppeers.crd.projectcalico.org

/registry/apiextensions.k8s.io/customresourcedefinitions/blockaffinities.crd.projectcalico.org

/registry/apiextensions.k8s.io/customresourcedefinitions/clusterinformations.crd.projectcalico.org

/registry/apiextensions.k8s.io/customresourcedefinitions/felixconfigurations.crd.projectcalico.org
......
```
##### 监听对象变化
```
etcdctl --endpoints https://localhost:2379 --cert /etc/kubernetes/pki/etcd/server.crt --key /etc/kubernetes/pki/etcd/server.key --cacert /etc/kubernetes/pki/etcd/ca.crt watch --prefix /registry/services/specs/default/mynginx

监听 mynginx 的变更
sh-5.0# < --key /etc/kubernetes/pki/etcd/server.key --cacert /etc/kubernetes/pki/etcd/ca.crt watch --prefix /registry/services/specs/default/mynginx
```

```
[root@bigdata-pro04 ~]# kubectl get svc
NAME         TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
kubernetes   ClusterIP   10.96.0.1    <none>        443/TCP   2d3h
```
变更 mynginx
```
[root@bigdata-pro04 ~]# kubectl edit svc mynginx
```
监听到对象变化
## APIServer
kube-APIServer 是 Kubernetes 最重要的核心组件之一，主要提供以下的功能：
1. 提供集群管理的 REST API 接口，包括：
1.1 认证 Authentication;
1.2 授权 Authorization;
1.3 准入 Admission（Mutating & Valiating）。
2. 提供其他模块之间的数据交互和通信的枢纽（其他模块通过 API Server 查询或修改数据，只有 API Server 才直接操作 etcd）。
3. APIServer 提供 etcd 数据缓存以减少集群对 etcd 的访问。

#### 演示 APIServer 的数据变化
##### watch
```
[root@bigdata-pro04 ~]# kubectl get svc -w
```
##### 查看 log
```
[root@bigdata-pro04 ~]# kubectl get svc -w -v 9
```
##### 变更服务信息
```
[root@bigdata-pro04 ~]# kubectl edit svc mynginx
```
监听到对象变化
## Controller Manager
- Controller Manager 是集群的大脑，是确保整个集群动起来的关键；
- 其作用是确保 Kubernetes 遵循声明式系统规范，确保系统的真实状态（Actual State）与用户定义的期望状态（Desired State 一直）；
- Controller Manager 是多个控制器的组合，每个 Controller 事实上都是一个 control loop，负责侦听其管控的对象，当对象发生变更时完成配置；
- Controller 配置失败通常会触发自动重试，整个集群会在控制器不断重试的机制下确保最终一致性（ Eventual Consistency）。
## Scheduler
特殊的 Controller，工作原理与其他控制器无差别；

Scheduler 的特殊职责在于监控当前集群所有未调度的Pod，并且获取当前集群所有节点的健康
状况和资源使用情况，为待调度 Pod 选择最佳计算节点，完成调度。

调度阶段分为：
- Predict：过滤不能满足业务需求的节点，如资源不足，端口冲突等。
- Priority：按既定要素将满足调度需求的节点评分，选择最佳节点。
- Bind：将计算节点与 Pod 绑定，完成调度。
## Kubelet
Kubernetes 的初始化系统（init system） 
- 从不同源获取 Pod 清单，并按需求启停 Pod 的核心组件：
  - Pod 清单可从本地文件目录，给定的 HTTPServer 或 KubeAPIServer 等源头获取；
  - Kubelet 将运行时，网络和存储抽象成了 CRI，CNI，CSI。 
- 负责汇报当前节点的资源信息和健康状态； 
- 负责 Pod 的健康检查和状态汇报。
## Kube-Proxy
- 监控集群中用户发布的服务，并完成负载均衡配置。
- 每个节点的Kube-Proxy都会配置相同的负载均衡策略，使得整个集群的服务发现建立在分布
式负载均衡器之上，服务调用无需经过额外的网络跳转（Network Hop）。
- 负载均衡配置基于不同插件实现：
  - userspace。 
  - 操作系统网络协议栈不同的 Hooks 点和插件：
    - iptables； 
    - ipvs。