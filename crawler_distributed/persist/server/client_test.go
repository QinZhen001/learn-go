package main

import (
	"learngo/crawler/engine"
	"learngo/crawler/model"
	"learngo/crawler_distributed/rpcsupport"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"

	// start ItemSaverServer
	go serveRpc(host, "test1")
	time.Sleep(time.Second)
	// start ItemSaverClient
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	// call save
	item := engine.Item{
		Url:  "http://album.zhenai.com/u/108906739",
		Id:   "108906739",
		Type: "zhenai",
		Payload: model.Profile{
			Name:       "安静的雪",
			Age:        34,
			Height:     162,
			Weight:     57,
			Income:     "3001-5000元",
			Gender:     "女",
			Marriage:   "离异",
			Education:  "大学本科",
			Occupation: "人事/行政",
			Hukou:      "山东菏泽",
			Xinzuo:     "牡羊座",
			House:      "已购房",
			Car:        "未购车",
		},
	}
	result := ""
	err = client.Call("ItemSaverService.Save", item, &result)
	if err != nil || result != "ok" {
		t.Errorf("result: %s; err: %s", result, err)
	}

}
