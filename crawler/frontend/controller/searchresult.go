package controller

import (
	"learngo/crawler/frontend/model"
	"learngo/crawler/frontend/view"
	"net/http"
	"strconv"
	"strings"

	"github.com/olivere/elastic"
)

type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

// loacalhost:8888/search?q=男 已购房&from=20
func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	q := strings.TrimSpace(req.FormValue("q"))
	form, err := strconv.Atoi(req.FormValue("from"))
	if err != nil {
		form = 0
	}

	// fmt.Fprintf(w, "q=%s, from=%d", q, form)
	var page model.SearchResult
	page = h.getSearchResult(q, form)
	err = h.view.Render(w, page)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

}

func CreateSearchResultHandler(template string) SearchResultHandler {
	client, err := elastic.NewClient(
		// Must turn off sniff in docker
		elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}

	return SearchResultHandler{
		view:   view.CreateSearchResultView(template),
		client: client,
	}
}
