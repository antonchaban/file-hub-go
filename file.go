package todo

type File struct {
	Id       int    `json:"file_id"`
	FileName string `json:"file_name"`
	FileDate string `json:"file_date"`
	FileSize int    `json:"file_size"`
	FilePath string `json:"file_path"`
}

type UsersFiles struct {
	Id     int
	UserId int
	FileId int
}
