package handlers

import (
	"day06/internal/models"
	"day06/internal/repositories"
	"html/template"
	"net/http"
	"strings"

	"github.com/google/uuid"
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

	formattedDescription := strings.ReplaceAll(string(htmlDescription), "\n", "<br>")

	dtoWithMD := struct {
		Id               string
		ShortDescription string
		Description      template.HTML
	}{
		Id:               post.Id,
		ShortDescription: post.ShortDescription,
		Description:      template.HTML(formattedDescription),
	}

	tmpl := template.Must(template.ParseFiles("../../web/templates/post_detail.html"))
	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, dtoWithMD)
}

func (r *PostHandler) PostCreate(w http.ResponseWriter, req *http.Request) {
	newPost := models.Post{
		Id:               uuid.New().String(),
		ShortDescription: req.FormValue("shortDescription"),
		Description:      req.FormValue("description"),
	}

	if _, err := r.repo.Create(&newPost); err != nil {
		http.Error(w, "Failed to create a post", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, req, "/post/" + newPost.Id, http.StatusMovedPermanently)
}

func (r *PostHandler) PostDelete(w http.ResponseWriter, req *http.Request) {
	uuid := req.FormValue("postId")

	model, err := r.repo.Get(uuid)
	if err != nil || model == nil {
		http.Error(w, "Not found with this uuid", http.StatusNotFound)
		return
	}
	if err = r.repo.Delete(uuid); err != nil {
		http.Error(w, "Failed to delete", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, req, "/page/1", http.StatusMovedPermanently)
}

func (r *PostHandler) PostUpdate(w http.ResponseWriter, req *http.Request) {
	uuid := req.FormValue("postId")

	model, err := r.repo.Get(uuid)
	if err != nil || model == nil {
		http.Error(w, "Not found with this uuid", http.StatusNotFound)
		return
	}

	model.Description = req.FormValue("description")
	model.ShortDescription = req.FormValue("shortDescription")

	if err = r.repo.Update(model); err != nil {
		http.Error(w, "Failed to update", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, req, "/post/" + model.Id, http.StatusMovedPermanently)
}