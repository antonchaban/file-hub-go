package service

import (
	fhub "github.com/antonchaban/file-hub-go"
	"github.com/antonchaban/file-hub-go/pkg/repository"
)

type Authorization interface {
	CreateUser(user fhub.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Folder interface {
	CreateFolder(userId int, folder fhub.Folder) (int, error)
	GetAllFolders(userId int) ([]fhub.Folder, error)
	GetById(userId, folderId int) (fhub.Folder, error)
	DeleteFolder(userId, folderId int) error
	UpdateFolder(userId, folderId int, input fhub.UpdateFolderInput) error
}

type File interface {
	CreateFile(userId, folderId int, file fhub.File) (int, error)
	GetAllFiles(userId, folderId int) ([]fhub.File, error)
	GetFileById(userId, fileId int) (fhub.File, error)
	DeleteFile(userId, fileId int) error
	UpdateFile(userId, fileId int, input fhub.UpdateFileInput) error
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
		File:          NewFileService(repos.File, repos.Folder),
	}
}
