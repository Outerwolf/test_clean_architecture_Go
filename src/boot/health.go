package boot

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterHealthCheck(mux *gin.Engine) {
	mux.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
}
