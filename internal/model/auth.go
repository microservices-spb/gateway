package model

type RequestData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Id       int64  `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
}

type ResponseData struct {
	Token string `json:"token"`
}
