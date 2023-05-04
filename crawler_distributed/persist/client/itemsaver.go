package client

import (
	"learngo/crawler/engine"
	"learngo/crawler_distributed/config"
	"learngo/crawler_distributed/rpcsupport"
	"log"
)

func ItemSaver(host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		for {
			item := <-out
			// Call RPC to save item
			result := ""
			client.Call(config.ElasticIndex, item, &result)
			if err != nil {
				log.Printf("Item %v save error: %v", item, err)
				continue
			}
		}
	}()
	return out, nil
}
