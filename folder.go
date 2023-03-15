package todo

type Folder struct {
	Id         int    `json:"id"`
	FolderName string `json:"folder_name"`
	FolderDate string `json:"folder_date"`
}

type UsersFolders struct {
	Id       int
	UserId   int
	FolderId int
}
