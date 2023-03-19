package fhub

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
