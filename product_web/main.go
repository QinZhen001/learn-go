package main

import (
	"learngo/internal"
	"learngo/internal/register"
	"learngo/util"

	uuid "github.com/satori/go.uuid"
)

var (
	consulRegistry register.ConsulRegistry
	randomId       string
)

// TODO: this
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

func main() {}
