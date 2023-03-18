package todo

import "errors"

type Folder struct {
	Id         int    `json:"id" db:"id"`
	FolderName string `json:"folder_name" db:"folder_name" binding:"required"`
	FolderDate string `json:"folder_date" db:"folderdate"`
}

type UsersFolders struct {
	Id       int
	UserId   int
	FolderId int
}

type UpdateFolderInput struct {
	FolderName *string `json:"folder_name"` // if empty - got nil
}

func (i UpdateFolderInput) Validate() error {
	if i.FolderName == nil {
		return errors.New("folder name is empty")
	}
	return nil
}
