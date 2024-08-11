package repositories

import (
	"day06/internal/models"
	"log/slog"
)

type Repository interface {
	Create(*models.Post, *slog.Logger) (uuid *string, err error)
	Get(uuid string, log *slog.Logger) (*models.Post, error)
	Update(*models.Post, *slog.Logger) error
	Delete(uuid string, log *slog.Logger) error
}
