package service

import (
	fhub "github.com/antonchaban/file-hub-go"
	"github.com/antonchaban/file-hub-go/pkg/repository"
)

type FileService struct {
	repo       repository.File
	folderRepo repository.Folder
}

func NewFileService(repo repository.File, folderRepo repository.Folder) *FileService {
	return &FileService{repo: repo, folderRepo: folderRepo}
}

func (s *FileService) CreateFile(userId, folderId int, file fhub.File) (int, error) {
	_, err := s.folderRepo.GetById(userId, folderId)
	if err != nil { // not exists or not owner
		return 0, err
	}

	return s.repo.CreateFile(folderId, file)
}

func (s *FileService) GetAllFiles(userId, folderId int) ([]fhub.File, error) {
	return s.repo.GetAllFiles(userId, folderId)
}

func (s *FileService) GetFileById(userId, fileId int) (fhub.File, error) {
	return s.repo.GetFileById(userId, fileId)
}

func (s *FileService) DeleteFile(userId, fileId int) error {
	return s.repo.DeleteFile(userId, fileId)
}

func (s *FileService) UpdateFile(userId, fileId int, input fhub.UpdateFileInput) error {
	return s.repo.UpdateFile(userId, fileId, input)
}
