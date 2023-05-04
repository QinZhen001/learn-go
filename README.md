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





## log 



### Fatal

[https://stackoverflow.com/questions/33885235/should-a-go-package-ever-use-log-fatal-and-when](https://stackoverflow.com/questions/33885235/should-a-go-package-ever-use-log-fatal-and-when)





## formatting

[https://gobyexample.com/string-formatting](https://gobyexample.com/string-formatting)



### %v

```
%v	the value in a default format
	when printing structs, the plus flag (%+v) adds field names
	
// 	%+v  打印结构体的时候，会把里面字段名也打印出来
```



### %s

For basic string printing use `%s`.

可以用来打印 byte[] 



## error

[https://www.cnblogs.com/qcrao-2018/p/11538387.html](https://www.cnblogs.com/qcrao-2018/p/11538387.html)





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
[https://www.cnblogs.com/liyutian/p/10050320.html](https://www.cnblogs.com/liyutian/p/10050320.html)
表达式 new(Type) 和 &Type{} 是等价的。

new 关键字时 t := new(T)，变量 t 则是一个指向 T 的指针



## don't close response.Body?

[https://stackoverflow.com/questions/33238518/what-could-happen-if-i-dont-close-response-body](https://stackoverflow.com/questions/33238518/what-could-happen-if-i-dont-close-response-body)





## string和[]byte

[https://zboya.github.io/post/golang_byte_slice_and_string/](https://zboya.github.io/post/golang_byte_slice_and_string/)



string:

```go
type stringStruct struct {
    str unsafe.Pointer
    len int
}
```

简单的来说字符串是一系列8位字节的集合，通常但不一定代表UTF-8编码的文本。字符串可以为空，但不能为nil。而且字符串的值是不能改变的。

[]byte:

```go
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}
```

array是数组的指针，len表示长度，cap表示容量。除了cap，其他看起来和string的结构很像。



标准转换

```go
// string to []byte
s1 := "hello"
b := []byte(s1)

// []byte to string
s2 := string(b)
```



### []byte转int

```go
var b = []byte{49,49} 
strconv.Atoi(string(b))
```







## dereference the pointer

对指针解引用





# docker



docker run -it alpine sh





docker run -d -p 80:80 nginx    

-d   后台一直运行 daemon

 -p 映射到物理机器的80端口





docker ps 查看进程





docker kill 973322dc1872

973322dc1872 => CONTAINER ID



## Daemon

Docker Daemon是**Docker架构中运行在后台的守护进程**，大致可以分为Docker Server、Engine和Job三部分。 Docker Daemon可以认为是通过Docker Server模块接受Docker Client的请求，并在Engine中处理请求，然后根据请求类型，创建出指定的Job并运行



# shell 编程

[https://wangdoc.com/bash/intro](https://wangdoc.com/bash/intro)



# elasticsearch



```tsx
 elastic.NewClient(
    // Must turn off sniff in docker
    // 不会自动转换地址
    elastic.SetSniff(false))
```



## Sniffing

[https://juejin.cn/post/6959746724995072031](https://juejin.cn/post/6959746724995072031)
