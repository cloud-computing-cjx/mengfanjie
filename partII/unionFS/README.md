# OverlayFS 文件系统练习
```
➜  test ll
总用量 0
➜  test mkdir upper lower merged work
➜  test echo "from lower" > lower/in_lower.txt
echo "from upper" > upper/in_upper.txt
echo "from lower" > lower/in_both.txt
echo "from upper" > upper/in_both.txt
➜  test ll
总用量 16K
drwxrwxr-x 2 cjx cjx 4.0K 11月  4 23:30 lower
drwxrwxr-x 2 cjx cjx 4.0K 11月  4 23:30 merged
drwxrwxr-x 2 cjx cjx 4.0K 11月  4 23:30 upper
drwxrwxr-x 2 cjx cjx 4.0K 11月  4 23:30 work
➜  test tree
.
├── lower
│   ├── in_both.txt
│   └── in_lower.txt
├── merged
├── upper
│   ├── in_both.txt
│   └── in_upper.txt
└── work

➜  test sudo mount -t overlay overlay -o lowerdir=`pwd`/lower,upperdir=`pwd`/upper,workdir=`pwd`/work `pwd`/merged
➜  test tree
.
├── lower
│   ├── in_both.txt
│   └── in_lower.txt
├── merged
│   ├── in_both.txt
│   ├── in_lower.txt
│   └── in_upper.txt
├── upper
│   ├── in_both.txt
│   └── in_upper.txt
└── work
    └── work [error opening dir]

5 directories, 7 files
➜  test cd merged 
➜  merged ll
总用量 12K
-rw-rw-r-- 1 cjx cjx 11 11月  4 23:30 in_both.txt
-rw-rw-r-- 1 cjx cjx 11 11月  4 23:30 in_lower.txt
-rw-rw-r-- 1 cjx cjx 11 11月  4 23:30 in_upper.txt
➜  merged cat in_both.txt 
from upper
➜  merged cat in_lower.txt 
from lower
➜  merged cat in_upper.txt 
from upper

```
# Docker 引擎架构
```
➜  ~ docker ps
CONTAINER ID   IMAGE          COMMAND                  CREATED          STATUS          PORTS     NAMES
24700075eb26   nginx:latest   "/docker-entrypoint.…"   20 minutes ago   Up 20 minutes   80/tcp    nginx
➜  ~ docker inspect nginx|grep Pid 
            "Pid": 23675,
            "PidMode": "",
            "PidsLimit": null,
➜  ~ ps -ef|grep 23675
root     23675 23652  0 23:24 ?        00:00:00 nginx: master process nginx -g daemon off;
systemd+ 23765 23675  0 23:24 ?        00:00:00 nginx: worker process
systemd+ 23766 23675  0 23:24 ?        00:00:00 nginx: worker process
systemd+ 23767 23675  0 23:24 ?        00:00:00 nginx: worker process
systemd+ 23768 23675  0 23:24 ?        00:00:00 nginx: worker process
systemd+ 23769 23675  0 23:24 ?        00:00:00 nginx: worker process
systemd+ 23771 23675  0 23:24 ?        00:00:00 nginx: worker process
systemd+ 23772 23675  0 23:24 ?        00:00:00 nginx: worker process
systemd+ 23774 23675  0 23:24 ?        00:00:00 nginx: worker process
systemd+ 23775 23675  0 23:24 ?        00:00:00 nginx: worker process
systemd+ 23776 23675  0 23:24 ?        00:00:00 nginx: worker process
systemd+ 23777 23675  0 23:24 ?        00:00:00 nginx: worker process
systemd+ 23778 23675  0 23:24 ?        00:00:00 nginx: worker process
systemd+ 23779 23675  0 23:24 ?        00:00:00 nginx: worker process
systemd+ 23780 23675  0 23:24 ?        00:00:00 nginx: worker process
systemd+ 23781 23675  0 23:24 ?        00:00:00 nginx: worker process
systemd+ 23782 23675  0 23:24 ?        00:00:00 nginx: worker process
systemd+ 23783 23675  0 23:24 ?        00:00:00 nginx: worker process
systemd+ 23784 23675  0 23:24 ?        00:00:00 nginx: worker process
systemd+ 23785 23675  0 23:24 ?        00:00:00 nginx: worker process
systemd+ 23786 23675  0 23:24 ?        00:00:00 nginx: worker process
cjx      27528 23789  0 23:45 pts/3    00:00:00 grep --color=auto --exclude-dir=.bzr --exclude-dir=CVS --exclude-dir=.git --exclude-dir=.hg --exclude-dir=.svn --exclude-dir=.idea --exclude-dir=.tox 23675
➜  ~ ps -ef|grep 23652
root     23652     1  0 23:24 ?        00:00:00 /usr/bin/containerd-shim-runc-v2 -namespace moby -id 24700075eb2615add6cc240a9227b218d70240d19790daba33f6f82924338e9f -address /run/containerd/containerd.sock
root     23675 23652  0 23:24 ?        00:00:00 nginx: master process nginx -g daemon off;
root     23870 23652  0 23:24 pts/0    00:00:00 sh
cjx      27543 23789  0 23:46 pts/3    00:00:00 grep --color=auto --exclude-dir=.bzr --exclude-dir=CVS --exclude-dir=.git --exclude-dir=.hg --exclude-dir=.svn --exclude-dir=.idea --exclude-dir=.tox 23652

```

