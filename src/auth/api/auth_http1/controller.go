package http1_auth

import "github.com/gin-gonic/gin"

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request SignUpRequest

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{
				"message": "Bad request",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "OK",
		})
	}
}

func SignIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	}
}

func SignOut() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	}
}

func RefreshToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	}
}
