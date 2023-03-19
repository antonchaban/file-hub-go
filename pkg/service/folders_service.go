package service

import (
	fhub "github.com/antonchaban/file-hub-go"
	"github.com/antonchaban/file-hub-go/pkg/repository"
)

type FolderService struct {
	repo repository.Folder
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

func (s *FolderService) Delete(userId, folderId int) error {
	return s.repo.Delete(userId, folderId)
}

func (s *FolderService) Update(userId, folderId int, input fhub.UpdateFolderInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, folderId, input)
}
