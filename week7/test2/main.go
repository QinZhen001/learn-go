package main

import (
	"time"

	"go.uber.org/zap"
)

func Newlogger() (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	config.OutputPaths = append(config.OutputPaths, "./7-2.log")
	return config.Build()
}

func main() {
	logger, err := Newlogger()
	if err != nil {
		panic(err)
	}
	sugar := logger.Sugar()
	defer sugar.Sync()
	e := "error"
	sugar.Info("又出错了",
		"err", e,
		"time", time.Now().Unix(),
	)
}
