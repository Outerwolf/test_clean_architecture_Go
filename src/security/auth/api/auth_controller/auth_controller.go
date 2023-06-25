package auth_controller

import (
	"github.com/auth/src/security/auth/domain/interfaces"
	"github.com/gin-gonic/gin"
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
	var request SignUpRequest

	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}
	request.Id = id

	ac.signUpUseCase.Execute(request)
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
