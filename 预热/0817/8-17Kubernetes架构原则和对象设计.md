# 作业
## 思考 stateful set 对象跟 deployment 对象的不同以及不同设计的目的。
类型特性 | Deployment | StatefulSet
---------- | ---------- | ---------- |
是否暴露到外网 | 可以 | 一般不
请求面向的对象 | serviceName | 指定pod的域名
灵活性 | 只能通过service/serviceIp访问到k8s自动转发的pod | 可以访问任意一个自定义的pod
易用性 | 只需要关心Service的信息即可 | 需要知道要访问的pod启动的名称、headlessService名称
PV/PVC绑定关系的稳定性（多replicas） | （pod挂掉后重启）无法保证初始的绑定关系 | 可以保证
pod名称稳定性 | 不稳定，因为是通过template创建，每次为了避免重复都会后缀一个随机数 | 稳定，每次都一样
启动顺序（多replicas） | 随机启动，如果pod宕掉重启，会自动分配一个node重新启动 | pod按 app-0、app-1...app-（n-1），如果pod宕掉重启，还会在之前的node上重新启动
停止顺序（多replicas） | 随机停止 | 倒序停止
集群内部服务发现  | 只能通过service访问到随机的pod | 可以打通pod之间的通信（主要是被发现）
性能开销 | 无需维护pod与node、pod与PVC 等关系 | 比deployment类型需要维护额外的关系信息

综上所述：
- 如果是不需额外数据依赖或者状态维护的部署，或者replicas是1，优先考虑使用Deployment；
- 如果单纯的要做数据持久化，防止pod宕掉重启数据丢失，那么使用pv/pvc就可以了；
- 如果要打通app之间的通信，而又不需要对外暴露，使用headlessService即可；
- 如果需要使用service的负载均衡，不要使用StatefulSet，尽量使用clusterIP类型，用serviceName做转发；
- 如果是有多replicas，且需要挂载多个pv且每个pv的数据是不同的，因为pod和pv之间是 一 一对应的，如果某个pod挂掉再重启，还需要连接之前的pv，不能连到别的pv上，考虑使用StatefulSet
- 能不用StatefulSet，就不要用

只能用StatefulSet:
最近在微软的aks平台上部署服务，由于Deployment在scale的时候需要动态申请volume，采取使用volumeClaimTemplates属性的方式来申请，当前Deployment对象（1.15）不支持这一属性，只有StatefulSet才有，因此不得不使用后者。目前看来有点本末倒置，不过不排除以后k8s会支持这一属性。

## 注意
如果使用StatefulSet，spec.serviceName需要指向headlessServiceName，且不能省略指定步骤，官方文档要求headlessService必须在创建StatefulSet之前创建完成，经过测试，如果没有也不会影响pod运行（pod还是running状态），只是不能拥有一个stable-network-id 集群内部不能访问到这个服务（如果这个服务不需要被发现，只需要去发现其他服务，则serviceName随便写一个也行），官方要求要在创建StatefulSet之前创建好headlessService，是为了让pod启动时能自动对应到service上。

之所以要指定一个headlessService，是因为admin可以给StatefulSet创建多个、多种类型的service，k8s不知道要用哪个service的名称当作集群内域名的一部分。

Deployment类型则不能有此参数，否则报错。