package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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

	// logger, err := NewLogger()
	// if err != nil {
	// 	panic(err)
	// }
	// defer logger.Sync()
	// sugar := logger.Sugar()
	// sugar.Info("failed to fetch URL ", "time:", 1)

	// ------------------  day 3  --------------------------
	// gin with zap
	r := gin.Default()
	zap.S()
	port := 9090
	zap.S().Info("starting server", zap.Int("port", port))
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	err := r.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		zap.S().Fatal("failed to start server", zap.Error(err))
	}
}
