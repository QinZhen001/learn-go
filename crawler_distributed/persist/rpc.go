package persist

import (
	"learngo/crawler/engine"

	"github.com/olivere/elastic"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaverService) Save(item engine.Item, result *string) error {

	return nil
}
