package persist

import (
	"context"
	"encoding/json"
	"learngo/crawler/engine"
	"learngo/crawler/model"
	"testing"

	"github.com/olivere/elastic"
)

func TestSave(t *testing.T) {
	expect := engine.Item{
		Url:  "http://album.zhenai.com/u/108906739",
		Id:   "108906739",
		Type: "zhenai",
		Payload: model.Profile{
			Name:       "安静的雪",
			Age:        34,
			Height:     162,
			Weight:     57,
			Income:     "3001-5000元",
			Gender:     "女",
			Marriage:   "离异",
			Education:  "大学本科",
			Occupation: "人事/行政",
			Hukou:      "山东菏泽",
			Xinzuo:     "牡羊座",
			House:      "已购房",
			Car:        "未购车",
		},
	}

	const index = "dating_test"
	// TODO:here use docker go client
	client, err := elastic.NewClient(
		elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}

	err = save(client, index, expect)

	if err != nil {
		panic(err)
	}

	resp, err := client.Get().Index("dating_profile").Type(expect.Type).Id(expect.Id).Do(context.Background())

	if err != nil {
		panic(err)
	}

	// t.Logf("%s", resp.Source)
	var actual engine.Item
	json.Unmarshal([]byte(*resp.Source), &actual)
	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	if actual != expect {
		t.Errorf("got %+v; expected %+v", actual, expect)
	}
}
