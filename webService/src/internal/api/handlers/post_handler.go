package handlers

import (
	"day06/internal/services"
	"net/http"
	"os"

	"github.com/russross/blackfriday/v2"
)



type PostHandler struct {
	service *services.PostService
}

func NewPostHandler(serv *services.PostService) *PostHandler {
	return &PostHandler{
		service: serv,
	}
}

func (r *PostHandler) PostGet(w http.ResponseWriter, req *http.Request) {

}

func (r *PostHandler) PostCreateForm(w http.ResponseWriter, req *http.Request) {
	form, err := os.ReadFile("../../../web/templates/PostCreateForm.md")
	if err != nil {
		http.Error(w, "Cannot open template createForm", http.StatusInternalServerError)
	}
	html := blackfriday.Run(form)

	w.Header().Set("Content-Type", "text/html")
	w.Write(html)
}

// POST Req
func (r *PostHandler) PostCreate(w http.ResponseWriter, req *http.Request) {
	
}