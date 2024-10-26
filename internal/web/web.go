package web

import (
	"github.com/djh20/gojuke/internal/config"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func Init() {
	if config.DebugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router = gin.Default()
	router.SetTrustedProxies(nil)

	registerApiRoutes()
}

func Run() {
	router.Run(config.ListenAddress)
}
