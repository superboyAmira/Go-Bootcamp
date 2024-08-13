package handlers

import (
	"day06/internal/repositories"
	"html/template"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/russross/blackfriday/v2"
)


type PostHandler struct {
	repo *repositories.PostRepository
}

func NewPostHandler(repository *repositories.PostRepository) *PostHandler {
	return &PostHandler{
		repo: repository,
	}
}

func (r *PostHandler) PostGet(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
    id := vars["id"]

	post, err := r.repo.Get(id)
    if err != nil {
        http.Error(w, "Post not found", http.StatusNotFound)
        return
    }
	htmlDescription := blackfriday.Run([]byte(post.Description))

	dtoWithMD := struct {
        Id              string
        ShortDescription string
        Description     template.HTML
    }{
        Id:              post.Id,
        ShortDescription: post.ShortDescription,
        Description:     template.HTML(htmlDescription),
    }

	tmpl := template.Must(template.ParseFiles("../../web/templates/post_detail.html"))
    w.Header().Set("Content-Type", "text/html")
    tmpl.Execute(w, dtoWithMD)
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