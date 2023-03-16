package repository

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

func NewRepository() *Repository {
	return &Repository{}
}
