# 一些基本命令
## build
编译(Go 语言不支持动态链接，因此编译时会将所有依赖编译进同一个二进制文件)
```
➜  helloworld git:(main) ✗ go build main.go
➜  helloworld git:(main) ✗ ./main 
os args is: [./main]
input parameter is: world
Hello world from Go
```
### 交叉编译(常用环境变量设置编译操作系统和 CPU 架构)
指定操作系统: GOOS=linux

GOOS=linux GOARCH=amd64 go build
```
➜  helloworld git:(main) ✗ GOOS=linux go build main.go
```
### 指定输出目录
go build –o bin/mybinary .
### 全支持列表
$GOROOT/src/go/build/syslist.go
## fmt
代码格式化
```
➜  helloworld git:(main) ✗ go fmt main.go 
```
## get
下载依赖
## install
编译安装
## mod
包管理
## run
运行
## test
测试
```
go test ./… -v 运行测试
go test命令扫描所有*_test.go为结尾的文件，惯例是将测试代码与正式代码放在同目录，
如 foo.go 的测试代码一般写在 foo_test.go
```
## vet
代码静态检查，发现可能的bug或者可疑的构造（检查编译器发现不了的错误）
### Print-format 错误，检查类型不匹配的print
```
str := “hello world!”
fmt.Printf("%d\n", str)
```
### Boolean 错误，检查一直为 true、false 或者冗余的表达式
```
fmt.Println(i != 0 || i != 1)
```
### Range 循环，比如如下代码主协程会先退出，go routine无法被执行
```
words := []string{"foo", "bar", "baz"} 
for _, word := range words {
    go func() {
        fmt.Println(word). 
    }()
}
```
### Unreachable的代码，如 return 之后的代码
### 其他错误，比如变量自赋值，error 检查滞后等
```
res, err := http.Get("https://www.spreadsheetdb.io/") 
defer res.Body.Close() 
if err != nil {
    log.Fatal(err)
}
```