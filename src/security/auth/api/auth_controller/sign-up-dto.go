package auth_controller

type SignUpRequest struct {
	Id       string `json:"id"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}
