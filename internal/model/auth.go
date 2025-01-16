package model

type RequestData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ResponseData struct {
	Token string `json:"token"`
}
