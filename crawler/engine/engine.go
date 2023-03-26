package engine

import (
	"learngo/crawler/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		body, error := fetcher.Fetch(r.Url)
		log.Printf("Fetching %s", r.Url)
		if error != nil {
			log.Printf("Fetcher: error fetching url %s: %v", r.Url, error)
			continue
		}
		parseResult := r.ParserFunc(body)
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}
