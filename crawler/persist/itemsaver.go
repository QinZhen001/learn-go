package persist

import (
	"context"
	"errors"
	"learngo/crawler/engine"
	"log"

	"github.com/olivere/elastic"
)

func ItemSaver(index string) (chan engine.Item, error) {
	// save in elasticsearch
	// use http.Post()
	// or use elasticsearch client
	client, err := elastic.NewClient(
		// Must turn off sniff in docker
		// 不会自动转换地址
		elastic.SetSniff(false))
	out := make(chan engine.Item)
	if err != nil {
		return nil, err
	}
	go func() {
		for {
			item := <-out
			err := Save(client, index, item)
			if err != nil {
				log.Printf("Item %v save error: %v", item, err)
				continue
			}
		}
	}()
	return out, nil
}

func Save(client *elastic.Client, index string, item engine.Item) (err error) {
	if item.Type == "" {
		return errors.New("Must supply Type")
	}

	indexService := client.Index().
		Index(index).
		Type(item.Type).BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)
	}

	indexService.Do(context.Background())

	if err != nil {
		return err
	}

	return nil
}
