package view

import (
	"learngo/crawler/frontend/model"
	"os"
	"testing"
)

func TestSearchResultView_Render(t *testing.T) {

	view := CreateSearchResultView("template.html")

	page := model.SearchResult{}
	page.Hits = 123
	page.Start = 0

	err := view.Render(os.Stdout, page)

	if err != nil {
		panic(err)
	}
}
