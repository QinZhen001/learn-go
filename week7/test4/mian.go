package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type RedisConfig struct {
	RedisHost string `mapstructure:"redisHost"`
	RedisPort int32  `mapstructure:"redisPort"`
}

type ServerConfig struct {
	ServerName  string      `mapstructure:"serverName"`
	RedisConfig RedisConfig `mapstructuure:"redisConfig"`
}

func GetEnv(s string) int {
	viper.AutomaticEnv()
	return viper.GetInt(s)
}

func main() {
	v := viper.New()
	dev := GetEnv("DEV")
	configFilePath := "dev-config.yaml"
	if dev == 1 {
		configFilePath = "prod-config.yaml"
	}
	fmt.Println(configFilePath)
	v.SetConfigFile(configFilePath)
	v.ReadInConfig()
	var serverConfig ServerConfig
	err := v.Unmarshal(&serverConfig)
	if err != nil {
		panic(err)
	}
	fmt.Println(serverConfig.ServerName)
	fmt.Println(serverConfig.RedisConfig.RedisHost)
	fmt.Println(serverConfig.RedisConfig.RedisPort)

}
