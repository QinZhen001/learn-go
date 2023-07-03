# go example

[https://gobyexample.com/](https://gobyexample.com/)



# 值传递

一定要记住，在 Go 语言中，**函数的参数传递只有值传递**，而且传递的实参都是原始数据的一份拷贝。如果拷贝的内容是值类型的，那么在函数中就无法修改原始数据；如果拷贝的内容是指针（或者可以理解为引用类型 `map`、`chan` 等），那么就可以在函数中修改原始数据。



# bufio

Package bufio implements buffered I/O. It wraps an io.Reader or io.Writer object, creating another object (Reader or Writer) that also implements the interface but provides buffering and some help for textual I/O.





# goroutine

go中的goroutins并不是同时在运行。事实上，如果没有在代码中通过runtime.GOMAXPROCS(n) 其中n是整数，指定使用多核的话，goroutins都是在一个线程里的，它们之间通过不停的让出时间片轮流运行，达到类似同时运行的效果。



**当一个goroutine发生阻塞，Go会自动地把与该goroutine处于同一系统线程的其他goroutines转移到另一个系统线程上去，以使这些goroutines不阻塞**



# runtime.Gosched

用于让出CPU时间片。这就像跑接力赛，A跑了一会碰到代码runtime.Gosched()就把接力棒交给B了，A歇着了，B继续跑。



# channels

[https://www.sohamkamani.com/golang/channels/](https://www.sohamkamani.com/golang/channels/)



# sync.WaitGroup

[https://gobyexample.com/waitgroups](https://gobyexample.com/waitgroups)



# syscall.SIGINT和syscall.SIGTERM

在Go程序中，syscall.SIGINT和syscall.SIGTERM是两个特定的操作系统信号常量。

- syscall.SIGINT 代表 "终止信号"，通常由终端用户按下组合键（Ctrl+C）发送给正在运行的程序。它用于请求进程终止运行，通常用于优雅地停止程序。
- syscall.SIGTERM 代表 "终止信号"，通常由操作系统发送给正在运行的进程，用于通知进程终止运行。与SIGINT不同的是，SIGTERM信号不一定是有用户产生的，可以由操作系统或其他程序发送。

在Go程序中，可以通过导入`os`和`os/signal`包来处理这些信号，例如使用`signal.Notify`函数来捕获并响应信号的发送。



# qit

