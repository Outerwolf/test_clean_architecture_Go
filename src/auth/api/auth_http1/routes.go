package http1_auth

import "github.com/gin-gonic/gin"

func RegisterHttpEndpoints(mux *gin.Engine) {
	mux.POST("/api/auth/sign-up", SignUp())
}
