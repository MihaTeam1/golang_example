package user

import (
	"github.com/MihaTeam1/golang_example/internal/repository"
	"github.com/MihaTeam1/golang_example/internal/structs"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(user structs.User) (int, error) {
	return s.repo.User.Create(user)
}

func (s *Service) GetById(id int) (structs.User, error) {
	return s.repo.User.GetById(id)
}

func (s *Service) GetByUsername(username string) (structs.User, error) {
	return s.repo.User.GetByUsername(username)
}
