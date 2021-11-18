# etcd 练习
启动新etcd集群
```
➜  ~ docker run --name etcd -d registry.aliyuncs.com/google_containers/etcd:3.5.0-0 /usr/local/bin/etcd
➜  ~ docker exec -it etcd sh
```
查看集群成员状态
```
➜  etcd-download-test ./etcdctl --endpoints=localhost:12379 member list --write-out=table
+------------------+---------+---------+------------------------+------------------------+------------+
|        ID        | STATUS  |  NAME   |       PEER ADDRS       |      CLIENT ADDRS      | IS LEARNER |
+------------------+---------+---------+------------------------+------------------------+------------+
| c9ac9fc89eae9cf7 | started | default | http://localhost:12380 | http://localhost:12379 |      false |
+------------------+---------+---------+------------------------+------------------------+------------+
```

```
sh-5.0# etcdctl put x 0
OK
sh-5.0# etcdctl get x
x
0
sh-5.0# etcdctl get x -w=json
{"header":{"cluster_id":14841639068965178418,"member_id":10276657743932975437,"revision":2,"raft_term":2},"kvs":[{"key":"eA==","create_revision":2,"mod_revision":2,"version":1,"value":"MA=="}],"count":1}

sh-5.0# etcdctl put x 1
OK
sh-5.0# etcdctl get x
x
1
sh-5.0# etcdctl get x --rev=2
x
0
```
