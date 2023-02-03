package Model

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (req LoginRequest) IsValid() bool {
	return req.Username != "" || req.Password != ""
}

type LoginResponse struct {
	Token string `json:"token"`
}
