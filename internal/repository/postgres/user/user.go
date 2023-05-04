package user

import (
	"github.com/MihaTeam1/golang_example/internal/structs"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(user structs.User) (int, error) {
	return 1, nil
}

func (r *Repository) GetById(id int) (structs.User, error) {
	return structs.User{}, nil
}

func (r *Repository) GetByUsername(username string) (structs.User, error) {
	return structs.User{}, nil
}
