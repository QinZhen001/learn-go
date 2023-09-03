package internal

import (
	"context"
	"fmt"
	"learngo/model"

	"github.com/olivere/elastic/v7"
)

var ESClient *elastic.Client

type EsConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

func InitEs() {
	host := fmt.Sprintf("%s:%d", AppConf.EsConfig.Host, AppConf.EsConfig.Port)
	var err error
	// https://github.com/olivere/elastic/issues/57
	// SetSniff false when use docker-compose
	ESClient, err = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(host))
	if err != nil {
		panic(err)
	}
	ok, err := ESClient.IndexExists(model.GetIndex()).Do(context.Background())
	if err != nil {
		panic(err)
	}
	if !ok {
		_, err := ESClient.CreateIndex(model.GetIndex()).BodyString(model.GetMapping()).Do(context.Background())
		if err != nil {
			panic(err)
		}
	}
}
