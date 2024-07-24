package service

import (
	"goday03/src/internal/app/model"
	"goday03/src/internal/app/repository"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
)

type PlaceService struct {
	reositoryPlaces *repository.PlaceRepository
}

func NewPlaceService(repository *repository.PlaceRepository) *PlaceService {
	return &PlaceService{reositoryPlaces: repository}
}

func (s *PlaceService) StorePageHandler(w http.ResponseWriter, req *http.Request) {
	pageStr := req.URL.Query().Get("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		http.Error(w, "Invalid page number", http.StatusBadRequest)
		return
	}
	limit := 10
	offset := (page - 1) * limit
	places, total, err := s.reositoryPlaces.GetPlaces(limit, offset)
	if err != nil {
		http.Error(w, "Invalid page", http.StatusBadRequest)
		return
	}
	tmplPath := filepath.Join("..", "..", "internal", "web", "html", "page.html")
	tmpl, err := template.ParseFiles(tmplPath)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Places     []model.Place
		Page       int
		TotalPages int
		PrevPage   int
		NextPage   int
	}{
		Places:     places,
		Page:       page,
		TotalPages: (total + limit - 1) / limit,
		PrevPage:   page - 1,
		NextPage:   page + 1,
	}

	tmpl.Execute(w, data)
}

func (s *PlaceService) BigStorePageHandler(w http.ResponseWriter, req *http.Request) {
	pageStr := req.URL.Query().Get("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		http.Error(w, "Invalid page number", http.StatusBadRequest)
		return
	}

	limit := 10
	batchSize := limit
	offset := (page - 1) * limit

	if offset > 13650 || offset < 0 {
		http.Error(w, "Invalid page number!", http.StatusBadRequest)
		log.Println(http.StatusBadRequest)
		return
	}

	allPlaces, err := s.reositoryPlaces.ScrollApiPlaces(batchSize)
	if err != nil {
		http.Error(w, "Error fetching places: "+err.Error(), http.StatusBadRequest)
		return
	}
	start := offset
	end := offset + limit
	if start > len(allPlaces) {
		start = len(allPlaces)
	}
	if end > len(allPlaces) {
		end = len(allPlaces)
	}
	places := allPlaces[start:end]
	total := len(allPlaces)

	tmplPath := filepath.Join("..", "..", "internal", "web", "html", "page.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Places     []model.Place
		Page       int
		TotalPages int
		PrevPage   int
		NextPage   int
	}{
		Places:     places,
		Page:       page,
		TotalPages: (total + limit - 1) / limit,
		PrevPage:   page - 1,
		NextPage:   page + 1,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}
