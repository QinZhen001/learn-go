package main

import (
	"learngo/crawler/engine"
	"learngo/crawler/scheduler"
	"learngo/crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		WorkerCount: 100,
		// Scheduler:     &scheduler.QueuedScheduler{},
		Scheduler: new(scheduler.SimpleScheduler),
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
