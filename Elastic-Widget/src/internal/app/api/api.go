package api

import (
	"net/http"
	"strconv"

	"github.com/olivere/elastic/v7"
)

type Api struct {
	client *elastic.Client
}

func NewApi(client *elastic.Client) (api *Api) {
	return &Api{client: client}
}

func (a *Api) GetPlacesApiHandler(w http.ResponseWriter, req *http.Request) {
	pageStr := req.URL.Query().Get("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		http.Error(w, "Cannot parse page number!", http.StatusBadRequest)
		return
	}
	if page > 13650 || page < 0 {
		http.Error(w, "Invalid page number!", http.StatusBadRequest)
	}
	
}