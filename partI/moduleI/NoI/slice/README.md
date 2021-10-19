# Slice
```
➜  slice git:(main) ✗ go build main.go 
➜  slice git:(main) ✗ ll
总用量 1.8M
-rwxrwxr-x 1 cjx cjx 1.7M 10月 19 22:25 main
-rw-rw-r-- 1 cjx cjx  442 10月 19 22:25 main.go
➜  slice git:(main) ✗ ./main 
mySlice [2 3]
fullSlice [1 2 3 4 5]
remove3rdItem [1 2 4 5]
```
## Make 和 New
- New 返回指针地址
- Make 返回第一个元素，可预设内存空间，避免未来的内存拷贝