package main

import (
	"learngo/crawler/engine"
	"learngo/crawler/persist"
	"learngo/crawler/scheduler"
	"learngo/crawler/zhenai/parser"
)

func main() {
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		WorkerCount: 10,
		Scheduler:   &scheduler.QueuedScheduler{},
		// Scheduler: new(scheduler.SimpleScheduler),
		ItemChan: itemChan,
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

	// e.Run(engine.Request{
	// 	Url:        "http://www.zhenai.com/zhenghun/shanghai",
	// 	ParserFunc: parser.ParseCity,
	// })
}
