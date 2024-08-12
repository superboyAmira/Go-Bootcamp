package services

import (
	"day06/internal/models"
	"day06/internal/repositories"
	"log/slog"
)

type PostService struct {
	reposiry repositories.Repository
}

func NewPostService(repo *repositories.PostRepository) *PostService {
	return &PostService{
		reposiry: repo,
	}
}

func (r *PostService) CreteObject(model *models.Post, log *slog.Logger) (*string, error) {
	uuid, err := r.reposiry.Create(model, log)
	return uuid, err
}
func (r *PostService) GetObject(uuid string, log *slog.Logger) (any, error) {
	model, err := r.reposiry.Get(uuid, log)
	return model, err
}

func (r *PostService) UpdateObject(model *models.Post, log *slog.Logger) error {
	err := r.reposiry.Update(model, log)
	return err
}

func (r *PostService) DeleteObject(uuid string, log *slog.Logger) error {
	err := r.reposiry.Delete(uuid, log)
	return err
}
