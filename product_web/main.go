package main

import (
	"fmt"
	"learngo/internal"
	"learngo/internal/register"
	"learngo/product_web/handler"
	"learngo/util"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

var (
	consulRegistry register.ConsulRegistry
	randomId       string
)

func init() {
	randomPort := util.GenRandomPort()
	if !internal.AppConf.Debug {
		internal.AppConf.ProductWebConfig.Port = randomPort
	}
	randomId = uuid.NewV4().String()
	consulRegistry = register.NewConsulRegistry(
		internal.AppConf.ConsulConfig.Host,
		internal.AppConf.ConsulConfig.Port,
	)
	consulRegistry.Register(
		internal.AppConf.ProductWebConfig.SrvName,
		randomId,
		internal.AppConf.ProductWebConfig.Port,
		internal.AppConf.ProductWebConfig.Tags,
	)
}

func main() {
	ip := internal.AppConf.ProductWebConfig.Host
	port := util.GenRandomPort()
	if internal.AppConf.Debug {
		port = internal.AppConf.ProductWebConfig.Port
	}
	addr := fmt.Sprintf("%s:%d", ip, port)
	r := gin.Default()

	productGroup := r.Group("/v1/product")
	{
		productGroup.GET("/list", handler.ProductListHandler)
		productGroup.POST("/add", handler.AddHandler)
		productGroup.POST("/update", handler.UpdateHandler)
		productGroup.POST("/delete", handler.DelHandler)
		productGroup.GET("/detail/:id", handler.DetailHandler)
	}
	r.GET("/health", handler.HealthHandler)

	go func() {
		err := r.Run(addr)
		if err != nil {
			zap.S().Panic(addr + "启动失败" + err.Error())
		} else {
			zap.S().Info(addr + "启动成功")
		}
	}()

	q := make(chan os.Signal)
	signal.Notify(q, syscall.SIGINT, syscall.SIGTERM)
	<-q

	// 中断信号（SIGINT）或终止信号（SIGTERM）
	// 当 web 服务退出时，注销服务
	err := consulRegistry.DeRegister(randomId)
	if err != nil {
		zap.S().Panic("注销失败" + randomId + ":" + err.Error())
	} else {
		zap.S().Info("注销成功" + randomId)
	}
}
