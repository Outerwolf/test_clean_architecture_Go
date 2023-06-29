package boot

import (
	"github.com/auth/src/apps/config"
)

//func CreateHTTPGinServer() *gin.Engine {
//	// Create a new mux.
//	return gin.Default()
//}
//
//func StartHttpGinServer(configuration *config.Configuration, mux *gin.Engine) {
//	http.ListenAndServe(configuration.HttpServer.Port, mux)
//}

func StartHttpServer(cfg *config.Configuration, router *Router) {
	router.Run(":" + cfg.HttpServer.Port)
}
