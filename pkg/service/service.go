package service

import "github.com/antonchaban/file-hub-go/pkg/repository"

type Authorization interface {
}

type File interface {
}

type Folder interface {
}

type Service struct {
	Authorization
	File
	Folder
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
