package api

import (
	"encoding/json"
	"goday03/src/internal/app/model"
	"net/http"
)

type ResponseHTTP200 struct {
	Name   string        `json:"name"`
	Total  int           `json:"total"`
	Places []model.Place `json:"places"`
}

type ResponseHTTP400 struct {
	Err string `json:"error"`
}

func (r *ResponseHTTP200) SendResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(*r)
}

func (r *ResponseHTTP400) SendResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	json.NewEncoder(w).Encode(*r)
}
