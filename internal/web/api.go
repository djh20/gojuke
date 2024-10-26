package web

import (
	"net/http"

	"github.com/djh20/gojuke/internal/music"
	"github.com/gin-gonic/gin"
)

func registerApiRoutes() {
	router.POST("/api/queue", func(c *gin.Context) {
		json := queueAddRequest{}

		if c.Bind(&json) == nil {
			c.Status(http.StatusOK)

			track := music.Track{
				FileName: json.Track,
			}

			music.AddTrackToQueue(track)
		}
	})

	router.POST("/api/playback/skip", func(c *gin.Context) {
		music.Skip()
		c.Status(http.StatusOK)
	})
}
