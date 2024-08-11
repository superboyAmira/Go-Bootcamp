package services

import "day06/internal/repositories"

type PostService struct {
	reposiry *repositories.Repository
}