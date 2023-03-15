package todo

type File struct {
	Id       int    `json:"id"`
	FileName string `json:"file_name"`
	FileDate string `json:"file_date"`
	FileSize int    `json:"file_size"`
	FileDesc string `json:"file_desc"`
}

type UsersFiles struct {
	Id     int
	UserId int
	FileId int
}
