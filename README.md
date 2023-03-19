# learn-go
:art: learn go

没有 while 用 for 代替

## Printf 和 Sprintf

Printf : 只可以打印出格式化的字符串。只可以直接输出字符串类型的变量

```go
  fmt.Printf("%d",a)
```

Sprintf：格式化并返回一个字符串而不带任何输出。    

```go
   s := fmt.Sprintf("a %s", "string") fmt.Printf(s)
```

## GO111MODULE
[GO111MODULE 及 Go 模块](https://zhuanlan.zhihu.com/p/417246469)



## 环境变量
export GOPATH=/Users/qinzhen/go
export PATH=$GOPATH/bin


## 下载包
go install这个命令建议安装那些有主文件(main.go),下载下来是作为一个xxx.exe存在于bin目录下的
go get一般用来下载我们项目中调用某个库时使用的命令。

go get
gopm来获取无法下载的包

## 单元测试
 go test

覆盖率测试：
go test -coverprofile=c.out

以html形式查看覆盖率报告：
go tool cover -html=c.out 

性能测试：
go test -bench .

cpu耗时报告：
go test -bench . -cpuprofile cpu.out
查看：
go tool pprof cpu.out
让后：
web  (以svg形式查看)


## api文档
查看文档：
go doc 

将本地所有项目包括系统api生成文档：
godoc -http :6060
