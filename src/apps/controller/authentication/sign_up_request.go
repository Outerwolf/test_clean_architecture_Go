package authentication

type SignUpRequest struct {
	Id       string `json:"id" validate:"uuid4"`
	UserName string `json:"username" validate:"email"`
	Password string `json:"password" binding:"required,gte=8,lte=32"`
}
