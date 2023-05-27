package boot

import (
	"github.com/auth/src/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateHTTPGinServer() *gin.Engine {
	// Create a new mux.
	return gin.Default()
}

func StartHttpGinServer(configuration *config.Configuration, mux *gin.Engine) {
	http.ListenAndServe(configuration.HttpServer.Port, mux)
}
