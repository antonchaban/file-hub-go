package todo

type User struct {
	Id       int    `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
