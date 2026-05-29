package model

type ReqeustLoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResponseSuccessLoginUser struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}
