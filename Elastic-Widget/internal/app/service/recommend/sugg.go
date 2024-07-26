package recommend

import (
	"errors"
	"goday03/src/internal/app/repository"
	"goday03/src/internal/app/service/api"
	"goday03/src/internal/app/service/response"
	"net/http"
	"strconv"
)

type Recommendation struct {
	repository *repository.PlaceRepository
}

func NewRecommendation(rep *repository.PlaceRepository) *Recommendation {
	return &Recommendation{
		repository: rep,
	}
}

func (rec *Recommendation) GetRecommendations(w http.ResponseWriter, req *http.Request) {
	token := api.TokenJWT{}
	if err := token.Validate(w, req); !err {
		return
	}
	
	std400resp := func(err error) {
		resp := &response.ResponseHTTP400{
			Err: err.Error(),
		}
		resp.SendResponse(w)
	}

	latReq := req.URL.Query().Get("lat")
	lonReq := req.URL.Query().Get("lon")
	if latReq == "" || lonReq == "" {
		std400resp(errors.New("empty geo-data"))
	}
	lat, err := strconv.ParseFloat(latReq, 64)
	if err != nil {
		std400resp(errors.New("cannot parse geo-data"))
		return
	}
	lon, err := strconv.ParseFloat(lonReq, 64)
	if err != nil {
		std400resp(errors.New("cannot parse geo-data"))
		return
	}
	places, err := rec.repository.GetRecommendations(lat, lon)
	if err != nil {
		std400resp(err)
	}
	resp := &response.ResponseHTTP200{
		Name:   "Recommendation",
		Places: places,
	}
	resp.SendResponse(w)
}
