package persist

import (
	"context"
	"log"

	"github.com/olivere/elastic"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		for {
			item := <-out
			log.Printf("Got item %+v", item)
			save(item)
		}
	}()
	return out
}

func save(item interface{}) (id string, err error) {
	// save in elasticsearch
	// use http.Post()
	// or use elasticsearch client
	client, err := elastic.NewClient(
		// Must turn off sniff in docker
		// 不会自动转换地址
		elastic.SetSniff(false))

	if err != nil {
		return "", err
	}

	resp, err := client.Index().Index("dating_profile").Type("zhenai").BodyJson(item).Do(context.Background())

	if err != nil {
		return "", err
	}

	return resp.Id, nil
}
