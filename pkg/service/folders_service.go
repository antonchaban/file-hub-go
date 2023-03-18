package service

import (
	todo "github.com/antonchaban/file-hub-go"
	"github.com/antonchaban/file-hub-go/pkg/repository"
)

type FolderService struct {
	repo repository.Folder
}

func NewFolderService(repo repository.Folder) *FolderService {
	return &FolderService{repo: repo}
}

func (s *FolderService) CreateFolder(userId int, folder todo.Folder) (int, error) {
	return s.repo.CreateFolder(userId, folder)
}

func (s *FolderService) GetAllFolders(userId int) ([]todo.Folder, error) {
	return s.repo.GetAllFolders(userId)
}

func (s *FolderService) GetById(userId, id int) (todo.Folder, error) {
	return s.repo.GetById(userId, id)
}
