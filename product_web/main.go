package main

import (
	"learngo/internal"
	"learngo/util"

	uuid "github.com/satori/go.uuid"
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
}

func main() {}
