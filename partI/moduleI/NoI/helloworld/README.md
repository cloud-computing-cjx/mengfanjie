go语言的 main() 函数的参数通过 os.Args 获取。
```
➜  helloworld git:(main) ✗ ./main 
os args is: [./main]
input parameter is: world
Hello world from Go
```

```
➜  helloworld git:(main) ✗ ./main cjx
os args is: [./main cjx]
input parameter is: world
Hello world from Go
```
将 cjx 赋值给 name
```
➜  helloworld git:(main) ✗ ./main --name cjx
os args is: [./main --name cjx]
input parameter is: cjx
Hello cjx from Go
```
