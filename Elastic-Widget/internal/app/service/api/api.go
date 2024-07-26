package api

import (
	"errors"
	"goday03/src/internal/app/model"
	"goday03/src/internal/app/repository"
	"goday03/src/internal/app/service/response"
	"net/http"
	"strconv"
)

const (
	limit int    = 10
	index string = "places"
)

type ApiService struct {
	repository *repository.PlaceRepository
}

func NewApi(client *repository.PlaceRepository) (api *ApiService) {
	return &ApiService{repository: client}
}

func (a *ApiService) GetPlacesApiHandler(w http.ResponseWriter, req *http.Request) {
	std400resp := func(err error) {
		resp := &response.ResponseHTTP400{
			Err: err.Error(),
		}
		resp.SendResponse(w)
	}

	pageStr := req.URL.Query().Get("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		std400resp(err)
		return
	}

	offset := (page - 1) * limit

	cnt, err := repository.GetDocumentCnt(a.repository.Client, index)
	if err != nil {
		std400resp(err)
		return
	}

	if offset > cnt || offset < 0 {
		std400resp(errors.New("invalid page number"))
		return
	}

	allPlaces, err := a.repository.ScrollApiPlaces(limit)
	if err != nil {
		std400resp(err)
		return
	}
	var places []model.Place
	if (limit + offset) > cnt {
		places = allPlaces[offset:cnt]
	} else {
		places = allPlaces[offset:(limit + offset)]
	}

	resp := &response.ResponseHTTP200{
		Name:   index,
		Total:  cnt,
		Places: places,
	}
	resp.SendResponse(w)
}
