package persist

import (
	"context"
	"encoding/json"
	"learngo/crawler/model"
	"testing"

	"github.com/olivere/elastic/v6"
)

func TestSave(t *testing.T) {
	expect := model.Profile{
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
	}

	id, err := save(expect)
	if err != nil {
		panic(err)
	}

	client, err := elastic.NewClient(
		elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}

	resp, err := client.Get().Index("dating_profile").Type("zhenai").Id(id).Do(context.Background())

	if err != nil {
		panic(err)
	}

	// t.Logf("%s", resp.Source)
	var actual model.Profile
	err = json.Unmarshal([]byte(*resp.Source), &actual)

	if err != nil {
		panic(err)
	}

	if actual != expect {
		t.Errorf("got %v; expected %v", actual, expect)
	}
}
