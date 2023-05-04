package service

import (
	"github.com/MihaTeam1/golang_example/internal/repository"
	"github.com/MihaTeam1/golang_example/internal/service/user"
	"github.com/MihaTeam1/golang_example/internal/structs"
)

type User interface {
	Create(user structs.User) (id int, err error)
	GetById(id int) (user structs.User, err error)
	GetByUsername(username string) (user structs.User, err error)
}

type Service struct {
	User User
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		User: user.NewService(repo),
	}
}
