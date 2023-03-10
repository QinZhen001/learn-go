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

