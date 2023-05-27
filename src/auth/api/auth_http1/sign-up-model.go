package http1_auth

type SignUpRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}
