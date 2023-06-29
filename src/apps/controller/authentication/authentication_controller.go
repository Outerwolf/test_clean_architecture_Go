package authentication

import (
	"github.com/auth/src/context/security/authentication/domain/interfaces"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
	signUpUseCase interfaces.ISignUpUseCase[SignUpRequest]
}

func NewAuthController(SignUpUseCase interfaces.ISignUpUseCase[SignUpRequest]) *AuthController {
	return &AuthController{
		signUpUseCase: SignUpUseCase,
	}
}

func (ac *AuthController) SignUp(c *gin.Context) {

	var data SignUpRequest

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ac.signUpUseCase.Execute(data)

	c.JSON(200, gin.H{
		"message": "From Sign UP",
	})
}

func (ac *AuthController) SignIn(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "OK",
	})
}

func (ac *AuthController) SignOut(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "OK",
	})
}

func (ac *AuthController) RefreshToken(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "OK",
	})
}
