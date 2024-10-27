package web

import (
	"embed"
	"log"

	"github.com/djh20/openjukebox/internal/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func Init(frontend embed.FS) {
	if config.DebugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router = gin.Default()
	router.SetTrustedProxies(nil)
	router.Use(cors.Default())
	router.Use(static.Serve("/", static.EmbedFolder(frontend, "frontend/dist")))

	registerApiRoutes()
}

func Run() {
	log.Println("Listening on", config.ListenAddress)
	router.Run(config.ListenAddress)
}
