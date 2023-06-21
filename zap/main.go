package main

import (
	"go.uber.org/zap"
)

func NewLogger() (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	config.OutputPaths = append(config.OutputPaths, "zap.log")
	return config.Build()
}

func main() {
	// ------------------  day 1  --------------------------

	// pro, err := zap.NewProduction()
	// if err != nil {
	// 	panic(err)
	// }
	// defer pro.Sync()
	// suger := pro.Sugar()
	// e := "error"
	// suger.Infof("failed to fetch URL %s", e)
	// suger.Info("failed to fetch URL ", e, "time:", time.Now().Unix())

	// ------------------  day 2  --------------------------
	logger, err := NewLogger()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	sugar := logger.Sugar()
	sugar.Info("failed to fetch URL ", "time:", 1)
}
