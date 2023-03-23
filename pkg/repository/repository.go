package repository

import (
	fhub "github.com/antonchaban/file-hub-go"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user fhub.User) (int, error)
	GetUser(username, password string) (fhub.User, error)
	AddTokenToBlacklist(token string) (int, error)
	IsTokenInBlacklist(token string) (bool, error)
}
type Folder interface {
	CreateFolder(userId int, folder fhub.Folder) (int, error)
	GetAllFolders(userId int) ([]fhub.Folder, error)
	GetById(userId, id int) (fhub.Folder, error)
	DeleteFolder(userId, folderId int) error
	UpdateFolder(userId, folderId int, input fhub.UpdateFolderInput) error
}

type File interface {
	CreateFile(folderId int, file fhub.File) (int, error)
	GetAllFiles(userId, folderId int) ([]fhub.File, error)
	GetFileById(userId, fileId int) (fhub.File, error)
	DeleteFile(userId, fileId int) error
	UpdateFile(userId, fileId int, input fhub.UpdateFileInput) error
}

type Repository struct {
	Authorization
	File
	Folder
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Folder:        NewFolderPostgres(db),
		File:          NewFilePostgres(db),
	}
}
