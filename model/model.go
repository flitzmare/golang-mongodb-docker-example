package model

type User struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Token    string `json:"token"`
}

type ResponseResult struct {
	Error  interface{} `json:"error"`
	Result interface{} `json:"result"`
}

type Error struct {
	Field   *string `json:"field"`
	Message *string `json:"message"`
}
