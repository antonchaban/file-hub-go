package fhub

import "errors"

type File struct {
	Id       int    `json:"file_id" db:"id"`
	FileName string `json:"file_name" db:"file_name" binding:"required"`
	FileDate string `json:"file_date" db:"file_date"`
	FileSize int    `json:"file_size" db:"file_size"`
	FilePath string `json:"file_path" db:"file_path" binding:"required"`
}

type UsersFiles struct {
	Id     int
	UserId int
	FileId int
}

type UpdateFileInput struct {
	FileName *string `json:"file_name"`
	FilePath *string `json:"file_path"`
}

func (i UpdateFileInput) Validate() error {
	if i.FileName == nil && i.FilePath == nil {
		return errors.New("file name and file path are empty")
	}
	return nil
}
