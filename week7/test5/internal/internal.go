package internal

import (
	config "learngo/week7/test5/conf"

	"github.com/spf13/viper"
)

var ServerConfig config.ServerConfig

func init() {
	v := viper.New()
	v.SetConfigFile("conf/dev-config.yaml")
	v.ReadInConfig()
	v.Unmarshal(&ServerConfig)
}
