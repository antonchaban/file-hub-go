package service

import (
	todo "github.com/antonchaban/file-hub-go"
	"github.com/antonchaban/file-hub-go/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Folder interface {
	CreateFolder(userId int, folder todo.Folder) (int, error)
	GetAllFolders(userId int) ([]todo.Folder, error)
	GetById(userId, folderId int) (todo.Folder, error)
	Delete(userId, folderId int) error
	Update(userId, folderId int, input todo.UpdateFolderInput) error
}

type File interface {
}

type Service struct {
	Authorization
	File
	Folder
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Folder:        NewFolderService(repos.Folder),
	}
}
