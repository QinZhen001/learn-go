package main

import (
	"time"

	"go.uber.org/zap"
)

func main() {
	pro, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	// 最后记得调用Sync()方法，将缓冲区的日志追加到文件中
	defer pro.Sync()
	// 写入日志
	sugar := pro.Sugar()
	e := "error"
	sugar.Infof("error:%s", e)
	sugar.Info("又出错了",
		"err",
		e,
		"time:", time.Now().Unix(),
	)
	// zap.Newproduction.logger 是 suger 的4～ 10倍， suger 使用了 golang的反射
	pro.Info("loger info test...",
		zap.String("name err", e),
		zap.Int("price err", 123),
	)
}
