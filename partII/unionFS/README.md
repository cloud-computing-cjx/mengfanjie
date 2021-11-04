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