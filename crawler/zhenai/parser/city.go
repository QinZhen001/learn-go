package parser

import (
	"learngo/crawler/engine"
	"regexp"
)

// <a href="http://album.zhenai.com/u/1879578227" target="_blank">期待你的出现</a>
var profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
var cityUrlRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/shanghai/[^"]+)"`)

func ParseCity(contents []byte) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		url := string(m[1])
		name := string(m[2])
		result.Items = append(result.Items, engine.Item{
			Url:     url,
			Type:    "zhenai",
			Id:      name,
			Payload: name,
		})
		result.Requests = append(result.Requests, engine.Request{
			Url:        url,
			ParserFunc: profileParser(name, url),
		})
	}

	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}

	return result
}
