package worker

import "learngo/crawler/engine"

type SerializedParser struct {
	Name string
	Args interface{}
}

type CrawService struct {
}

func (CrawService) Process(req engine.Request, result *engine.ParseResult) error {
	//1.将req反序列化
	//2.调用真正的rpc服务
	//3.将结果反序列化
	return nil
}
