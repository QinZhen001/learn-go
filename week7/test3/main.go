package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type DBConfig struct {
	UserName string `mapstructure:"userName"`
	Password string `mapstructure:"password"`
}

func main() {
	v := viper.New()
	v.SetConfigFile("dev-config.yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println(v.Get("name"))

	var dbConfig DBConfig
	err = v.Unmarshal(&dbConfig)
	if err != nil {
		panic(err)
	}
	fmt.Println(dbConfig.UserName)
	fmt.Println(dbConfig.Password)

}
