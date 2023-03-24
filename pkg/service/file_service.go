package service

import (
	fhub "github.com/antonchaban/file-hub-go"
	"github.com/antonchaban/file-hub-go/pkg/repository"
)

type FileService struct {
	fileRepo   repository.File
	folderRepo repository.Folder
}

func NewFileService(fileRepo repository.File, folderRepo repository.Folder) *FileService {
	return &FileService{fileRepo: fileRepo, folderRepo: folderRepo}
}

func (s *FileService) CreateFile(userId, folderId int, file fhub.File) (int, error) {
	_, err := s.folderRepo.GetById(userId, folderId)
	if err != nil { // not exists or not owner
		return 0, err
	}

	return s.fileRepo.CreateFile(folderId, file)
}

func (s *FileService) GetAllFiles(userId, folderId int) ([]fhub.File, error) {
	return s.fileRepo.GetAllFiles(userId, folderId)
}

func (s *FileService) GetFileById(userId, fileId int) (fhub.File, error) {
	return s.fileRepo.GetFileById(userId, fileId)
}

func (s *FileService) DeleteFile(userId, fileId int) error {
	return s.fileRepo.DeleteFile(userId, fileId)
}

func (s *FileService) UpdateFile(userId, fileId int, input fhub.UpdateFileInput) error {
	return s.fileRepo.UpdateFile(userId, fileId, input)
}
