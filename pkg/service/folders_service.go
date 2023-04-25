package service

import (
	fhub "github.com/antonchaban/file-hub-go"
	"github.com/antonchaban/file-hub-go/pkg/repository"
)

type FolderService struct {
	repo repository.Folder
}

type Folder interface {
	CreateFolder(userId int, folder fhub.Folder) (int, error)
	GetAllFolders(userId int) ([]fhub.Folder, error)
	GetById(userId, folderId int) (fhub.Folder, error)
	DeleteFolder(userId, folderId int) error
	UpdateFolder(userId, folderId int, input fhub.UpdateFolderInput) error
}

func NewFolderService(repo repository.Folder) *FolderService {
	return &FolderService{repo: repo}
}

func (s *FolderService) CreateFolder(userId int, folder fhub.Folder) (int, error) {
	return s.repo.CreateFolder(userId, folder)
}

func (s *FolderService) GetAllFolders(userId int) ([]fhub.Folder, error) {
	return s.repo.GetAllFolders(userId)
}

func (s *FolderService) GetById(userId, folderId int) (fhub.Folder, error) {
	return s.repo.GetById(userId, folderId)
}

func (s *FolderService) DeleteFolder(userId, folderId int) error {
	return s.repo.DeleteFolder(userId, folderId)
}

func (s *FolderService) UpdateFolder(userId, folderId int, input fhub.UpdateFolderInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.UpdateFolder(userId, folderId, input)
}
