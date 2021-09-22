# 运行 Makefile 文件
```
➜  golang git:(main) ✗ cd httpserver 
➜  httpserver git:(main) ✗ make push
```
# Namespace(用来做隔离)
namespace类型 | 隔离资源
---- | ----
IPC | System V IPC 和 POSIX 消息队列
Network | 网络设备、网络协议栈、网络端口等
PID | 进程
Mount | 挂载点
UTS | 主机名和域名
USR | 用户和用户组
```
➜  ~ docker run -it centos
[root@edfe3a397e2d /]# ps -ef
UID        PID  PPID  C STIME TTY          TIME CMD
root         1     0  0 08:59 pts/0    00:00:00 /bin/bash
root        15     1  0 08:59 pts/0    00:00:00 ps -ef
[root@edfe3a397e2d /]# ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
23: eth0@if24: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default 
    link/ether 02:42:ac:11:00:02 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 172.17.0.2/16 brd 172.17.255.255 scope global eth0
       valid_lft forever preferred_lft forever
```
## namespace d的常用操作
### 查看当前系统的 namespace
```
lsns -t <type>
```
lsns 操作举例
```
➜  ~ lsns --help

用法：
 lsns [选项] [<名字空间>]

列出系统名字空间。

选项：
 -J, --json              使用 JSON 输出格式
 -l, --list             使用列表格式的输出
 -n, --noheadings       不打印标题
 -o, --output <list>    定义使用哪个输出列
 -p, --task <pid>       打印进程名字空间
 -r, --raw              使用原生输出格式
 -u, --notruncate       不截断列中的文本
 -t, --type <name>      名字空间类型(mnt, net, ipc, user, pid, uts, cgroup)

 -h, --help             display this help
 -V, --version          display version

Available output columns:
          NS  名字空间标识符 (inode 号)
        TYPE  名字空间类型
        PATH  名字空间路径
      NPROCS  名字空间中的进程数
         PID  名字空间中的最低 PID
        PPID  PID 的 PPID
     COMMAND  PID 的命令行
         UID  PID 的 UID
        USER  PID 的用户名

更多信息请参阅 lsns(8)。
➜  ~ lsns -t net
        NS TYPE NPROCS   PID USER COMMAND
4026531992 net     142  1193 cjx  /lib/systemd/systemd --user
4026532713 net      29  7216 cjx  /opt/google/chrome/chrome --type=renderer --field-trial-handle=2177232341976017818,13168316596007230620,131072 --disable-gpu-compositing --lang=zh-CN --enable-crashpad --crashpa
4026533274 net       1 23195 cjx  /opt/google/chrome/nacl_helper
4026534293 net       2 25021 cjx  /usr/share/atom/atom --type=zygote
```
### 查看某进程的 namaspace
```
ls -la /proc/<pid>/ns/
```
操作举例
```
➜  ~ docker ps|grep centos
edfe3a397e2d   centos              "/bin/bash"              8 minutes ago   Up 8 minutes             tender_gates
➜  ~ docker inspect edfe3a397e2d|grep -i pid
            "Pid": 16112,
            "PidMode": "",
            "PidsLimit": null,
➜  ~ nsenter -t 16112 -n ip a
nsenter: 打不开 /proc/16112/ns/net: 权限不够
➜  ~ sudo nsenter -t 16112 -n ip a
[sudo] cjx 的密码： 
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
23: eth0@if24: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default 
    link/ether 02:42:ac:11:00:02 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 172.17.0.2/16 brd 172.17.255.255 scope global eth0
       valid_lft forever preferred_lft forever

➜  ~ sudo nsenter -t 16112 -n netstat -na 
激活Internet连接 (服务器和已建立连接的)
Proto Recv-Q Send-Q Local Address           Foreign Address         State      
活跃的UNIX域套接字 (服务器和已建立连接的)
Proto RefCnt Flags       Type       State         I-Node   路径
```
### 进入某 namespace 运行命令
```
nsenter -t <pid> -n ip addr
```
## namespace练习
在新的 network namespace 执行 sleep 指令
```
➜  ~ sudo unshare -fn sleep 60
```
查看进程信息
```
➜  ~ ps  -ef|grep sleep
root     13444 18666  0 17:19 pts/7    00:00:00 sudo unshare -fn sleep 60
root     13445 13444  0 17:19 pts/7    00:00:00 unshare -fn sleep 60
root     13446 13445  0 17:19 pts/7    00:00:00 sleep 60
cjx      13820 13760  0 17:19 pts/8    00:00:00 grep --color=auto --exclude-dir=.bzr --exclude-dir=CVS --exclude-dir=.git --exclude-dir=.hg --exclude-dir=.svn --exclude-dir=.idea --exclude-dir=.tox sleep
```
查看网络 Namespace
```
➜  ~ sudo lsns -t net
4026534451 net       2 15900 root   unshare -fn sleep 60
```
进入该进程所在的 Namespace 查看网络配置
```
nsenter -t <pid> -n ip a
```
# Cgroups
- Cgroups （Control Groups）是 Linux 下用于对一个或一组进程进行资源控制和监控的机制；
- 可以对诸如 CPU 使用时间、内存、磁盘 I/O 等进程所需的资源进行限制；
- 不同资源的具体管理工作由相应的 Cgroup 子系统（Subsystem）来实现 ；
- 针对不同类型的资源限制，只要将限制策略在不同的的子系统上进行关联即可 ；
- Cgroups 在不同的系统资源管理子系统中以层级树（Hierarchy）的方式来组织管理：每个Cgroup 都可以包含其他的子 Cgroup，因此子 Cgroup 能使用的资源除了受本 Cgroup 配置的资源参数限制，还受到父 Cgroup 设置的资源限制 。
## 可配额/可度量 - Control Groups (cgroups)
cgroups 实现了对资源的配额和度量。
- blkio：这个子系统设置限制每个块设备的输入输出控制。例如:磁盘，光盘以及 USB 等等；
- cpu：这个子系统使用调度程序为 cgroup 任务提供 CPU 的访问；
- cpuacct：产生 cgroup 任务的 CPU 资源报告；
- cpuset：如果是多核心的CPU，这个子系统会为 cgroup 任务分配单独的 CPU 和内存；
- devices：允许或拒绝 cgroup 任务对设备的访问；
- freezer：暂停和恢复 cgroup 任务；
- memory：设置每个 cgroup 的内存限制以及产生内存资源报告；
- net_cls：标记每个网络包以供 cgroup 方便使用；
- ns：名称空间子系统；
- pid: 进程标识子系统
## CPU子系统
- cpu.shares：可出让的能获得 CPU 使用时间的相对值。
- cpu.cfs_period_us：cfs_period_us 用来配置时间周期长度，单位为 us（微秒）。
- cpu.cfs_quota_us：cfs_quota_us 用来配置当前 Cgroup 在 cfs_period_us 时间内最多能使用的 CPU时间数，单位为 us（微秒）。
- cpu.stat ：Cgroup 内的进程使用的 CPU 时间统计。
- nr_periods ：经过 cpu.cfs_period_us 的时间周期数量。
- nr_throttled ：在经过的周期内，有多少次因为进程在指定的时间周期内用光了配额时间而受到限制。
- throttled_time ：Cgroup 中的进程被限制使用 CPU 的总用时，单位是 ns（纳秒）。

```
➜  ~ cd /sys/fs/cgroup 
➜  cgroup ls
blkio  cpu  cpuacct  cpu,cpuacct  cpuset  devices  freezer  hugetlb  memory  net_cls  net_cls,net_prio  net_prio  perf_event  pids  rdma  systemd  unified
➜  cgroup cd cpu
➜  cpu ls
```
### CPU相对值
```
➜  cpu cat cpu.shares 
1024
```
cpu.cfs_quota_us ：能用的CPU时间片（-1表示不做限制）
cpu.cfs_period_us ：总的CPU时间片
```
➜  cpu cat cpu.cfs_quota_us cpu.cfs_period_us 
-1
100000
```
#### 测试
```
➜  cpu sudo mkdir cpudemo
[sudo] cjx 的密码： 
➜  cpu cd cpudemo 
➜  cpudemo ls
cgroup.clone_children  cpuacct.stat   cpuacct.usage_all     cpuacct.usage_percpu_sys   cpuacct.usage_sys   cpu.cfs_period_us  cpu.shares  cpu.uclamp.max  notify_on_release
cgroup.procs           cpuacct.usage  cpuacct.usage_percpu  cpuacct.usage_percpu_user  cpuacct.usage_user  cpu.cfs_quota_us   cpu.stat    cpu.uclamp.min  tasks
```
执行busyloop
```
➜  golang git:(main) ✗ examples/busyloop/busyloop 

显示 busyloop 消耗CPU在200左右
➜  cpudemo top

top - 18:01:08 up 5 days,  3:41,  1 user,  load average: 2.56, 1.53, 1.25
任务: 635 total,   4 running, 495 sleeping,   0 stopped,   0 zombie
%Cpu(s): 14.9 us,  0.9 sy,  0.0 ni, 82.8 id,  0.0 wa,  0.0 hi,  1.3 si,  0.0 st
KiB Mem : 13184452+total, 95782152 free, 20609144 used, 15453232 buff/cache
KiB Swap:  2097148 total,  2097148 free,        0 used. 10971814+avail Mem 

进�� USER      PR  NI    VIRT    RES    SHR �  %CPU %MEM     TIME+ COMMAND                                                                                                                                        
 9500 cjx       20   0  702420   1160    644 R 200.3  0.0   4:17.07 busyloop    
```
控制资源

添加要管理的进程号
```
bigdata# cd /sys/fs/cgroup/cpu/cpudemo 

bigdata# echo 9500 > cgroup.procs
```
限制可以使用的CPU资源
```
bigdata# echo 10000 > cpu.cfs_quota_us
```
查看测试结果（资源被限制在10左右）
```
bigdata# cat cpu.cfs_quota_us cpu.cfs_period_us 
10000
100000

bigdata# top

top - 18:04:50 up 5 days,  3:45,  1 user,  load average: 2.67, 2.12, 1.56
任务: 637 total,   2 running, 499 sleeping,   0 stopped,   0 zombie
%Cpu(s):  2.2 us,  1.4 sy,  0.0 ni, 96.1 id,  0.0 wa,  0.0 hi,  0.4 si,  0.0 st
KiB Mem : 13184452+total, 95768416 free, 20607660 used, 15468452 buff/cache
KiB Swap:  2097148 total,  2097148 free,        0 used. 10971758+avail Mem 

进�� USER      PR  NI    VIRT    RES    SHR �  %CPU %MEM     TIME+ COMMAND                                                                                                                                        
 9500 cjx       20   0  702420   1160    644 R  12.5  0.0  11:34.49 busyloop 
```
## cpuacct 子系统
用于统计 Cgroup 及其子 Cgroup 下进程的 CPU 的使用情况。
### cpuacct.usage
包含该 Cgroup 及其子 Cgroup 下进程使用 CPU 的时间，单位是 ns（纳秒）。
### cpuacct.stat
包含该 Cgroup 及其子 Cgroup 下进程使用的 CPU 时间，以及用户态和内核态的时间。
## memory 子系统
### memory.usage_in_bytes
cgroup下进程使用的内存，包含cgroup及其子cgroup下的进程使用的内存。
### memory.max_usage_in_bytes
cgroup下进程使用内存的最大值，包含子cgroup的内存使用量。
### memory.limit_in_bytes
设置Cgroup下进程最多能使用的内存。如果设置为-1，表示对该cgroup的内存使用不做限
制。
### memory.oom_control
设置是否在Cgroup中使用OOM（Out of Memory）Killer，默认为使用。当属于该cgroup
的进程使用的内存超过最大的限定值时，会立刻被OOM Killer处理。
### memory 子系统练习
## 文件系统
## Docker的文件系统
```
docker history <image id>
```