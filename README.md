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



# sql.NullString and *string

[https://stackoverflow.com/questions/40092155/difference-between-string-and-sql-nullstring](https://stackoverflow.com/questions/40092155/difference-between-string-and-sql-nullstring)

```tsx
type NullString struct {
    String string
    Valid  bool // Valid is true if String is not NULL
}
```

As you can see, `sql.NullString` is a way to represent null string coming from SQL (which correspond to "NULL"). On the other hand, a nil `*string` is a pointer to a string which is nil, so the two are different.



There's no effective difference. We thought people might want to use NullString because it is so common and perhaps expresses the intent more clearly than *string. But either will work

使用 sql.NullString 更清晰



# Gorm Associations

Gorm Associations 是 Gorm 框架中用来定义和管理数据库表之间关系的一种功能。

使用 Gorm Associations，可以简化数据库表之间的关联操作，比如定义一对一关系、一对多关系、多对多关系等。通过定义这些关系，可以方便地进行数据库查询和操作。

应用 Gorm Associations 的主要优点包括：

1. 简化代码：使用 Gorm Associations 可以简化关联关系的定义和查询操作，减少冗余代码。
2. 方便的查询：通过定义关联关系，可以方便地进行复杂的查询操作，包括跨表查询、嵌套查询等。
3. 数据库一致性：使用 Gorm Associations 可以确保数据库中的关联关系保持一致性，避免了手动管理关联关系带来的问题。
4. 代码可读性：通过使用 Gorm Associations，可以更清晰地表达实体之间的关系，提高代码可读性和可维护性。

总之，Gorm Associations 提供了一种方便、简化和规范的方式来管理数据库表之间的关系，提高了开发效率和代码质量。





# 其他



## protoc 使用

protoc go-grpc_out 是一个 protoc 的插件，用于生成 Go 的 gRPC 代码。以下是它的使用方法：

1. 确保已经安装了 protoc 编译器和 Go 的 protoc-gen-go 插件。
2. 定义一个 .proto 文件，其中包含了 gRPC 服务的接口定义和消息定义。
3. 使用以下命令生成 Go 的 gRPC 代码：

```bash
protoc book.proto --go_out=. --go-grpc_out=.
```

`--go_out=.` 表示生成的 Go 代码将保存在当前目录下，`--go-grpc_out=.` 表示生成的 gRPC 代码同样保存在当前目录下



## Mysql 无法启动

[Mac 下 Mysql 无法启动](https://juejin.cn/post/6844904088098832398)

```bash
sudo /usr/local/mysql/support-files/mysql.server start
```





