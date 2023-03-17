package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
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
	return &Repository{}
}
