package repository

import (
	todo "github.com/antonchaban/file-hub-go"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type File interface {
}

type Folder interface {
}

type Repository struct {
	Authorization
	File
	Folder
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
