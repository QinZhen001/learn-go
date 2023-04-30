package engine

import (
	"learngo/crawler/fetcher"
	"log"
)

func worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s", r.Url)
	body, error := fetcher.Fetch(r.Url)
	if error != nil {
		log.Printf("Fetcher: error fetching url %s: %v", r.Url, error)
		return ParseResult{}, error
	}
	parseResult := r.ParserFunc(body)
	return parseResult, nil
}
