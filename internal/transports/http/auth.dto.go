package http

type RegisterData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
