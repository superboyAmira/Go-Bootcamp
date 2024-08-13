package handlers

import (
	"day06/internal/models"
	"day06/internal/repositories"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
	// "github.com/russross/blackfriday/v2"
)

const (
	limitPosts = 3
)

type IndexHandler struct {
	repository *repositories.PostRepository
	sett PaginationData
}

type PaginationData struct {
	Posts []models.Post

	// pagination query settings
	Page int
	Offset int

	// pagination html settings
	NextPage int
	PrevPage int
	Next bool
	Prev bool
}

func NewIndexHandler(repo *repositories.PostRepository) *IndexHandler {
	return &IndexHandler{
		repository: repo,
		sett : PaginationData{
			Page: 1,
			Offset: 0,
			NextPage: 2,
			PrevPage: 0,
			Next: true,
			Prev: false,
		},
	}
}

func (r *IndexHandler) MainIndexHandler(w http.ResponseWriter, req *http.Request) {
	var err error
	vars := mux.Vars(req)
	r.sett.Page, err = strconv.Atoi(vars["pageNum"])

	if err != nil || r.sett.Page < 1 {
		r.sett.Page = 1
	}
	r.sett.Offset = (r.sett.Page - 1) * limitPosts

	r.sett.Posts, err = r.repository.GetAll(limitPosts, r.sett.Offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	count, err := r.repository.Count()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	r.sett.NextPage = r.sett.Page + 1
	r.sett.PrevPage = r.sett.Page - 1
	r.sett.Next = *count > int64(r.sett.Page * limitPosts) 
	r.sett.Prev = r.sett.Page > 1

	w.Header().Set("Content-Type", "text/html")
	tmpl := template.Must(template.ParseFiles("../../web/templates/index.html"))
	tmpl.Execute(w, r.sett)
}