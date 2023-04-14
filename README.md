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

## GOPATH 和 GOROOT
GOROOT：GOROOT就是Go的安装目录，（类似于java的JDK）
GOPATH：GOPATH是我们的工作空间,保存go项目代码和第三方依赖包

## 环境变量
export GOPATH=/Users/qinzhen/go
export PATH=$GOPATH/bin:/usr/local/go/bin
注意：go要多设置两个PATH

最终的PATH通过: 一起串联起来
export PATH=$PNPM_HOME:$PATH:/Users/qinzhen/go/bin:/usr/local/go/bin

查看环境变量：
echo $PATH


## 下载包
go install这个命令建议安装那些有主文件(main.go),下载下来是作为一个xxx.exe存在于bin目录下的
go get一般用来下载我们项目中调用某个库时使用的命令。

go get： 下载包 & 更新go.mod，不安装二进制
go install ：已下载的包，安装二进制

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
(在项目根路径执行)
godoc -http :6060



## run
go run命令会编译源码，并且直接执行源码的 main() 函数，不会在当前目录留下可执行文件。

go run xxx.go 
go run -race xxx.go 


## build
变成一个可执行文件，将始终打印出Hello World。如果我们想让程序再次运行，我们不需要再次编译程序，我们只需要运行可执行文件。


## tool1
pprof 分析性能
go tool pprof http://localhost:8888/debug/profile


## doc 
查看文档
godoc -http :8888


## io.Reader,io.ReadCloser和io.ReadSeeker
[https://github.com/wuyongxiu/wuyongxiu.github.io/issues/15](https://github.com/wuyongxiu/wuyongxiu.github.io/issues/15)


## Struct 和 Interface
Interface与struct类似，但只包含一些抽象方法。 在Go中，Interface定义了通用行为的抽象。

```tsx
//declare a rectangle struct
type rectangle struct {
    length int
    width  int
}

//declare an interface with area() as a member
type shape interface {
    area() int
}

//declare a method area()
//the rectangle struct implements area() method in shape interface
func (r rectangle) area() int {
    return r.length * r.width
}

//declare a method with type shape as a parameter
func info(s shape) {
    fmt.Println("the area: ", s.area())
}

```


# 补充

## don't close response.Body?
[https://stackoverflow.com/questions/33238518/what-could-happen-if-i-dont-close-response-body](https://stackoverflow.com/questions/33238518/what-could-happen-if-i-dont-close-response-body)
