package web

import (
	"net/http"

	"github.com/djh20/openjukebox/internal/music"
	"github.com/gin-gonic/gin"
)

func registerApiRoutes() {
	api := router.Group("/api")
	{
		api.GET("/tracks", handleTracksRequest)

		api.POST("/playback/queue", handleQueueRequest)
		api.POST("/playback/volume", handleVolumeRequest)
		api.POST("/playback/skip", handleSkipRequest)
	}
}

func handleTracksRequest(c *gin.Context) {
	c.JSON(http.StatusOK, music.Tracks)
}

func handleQueueRequest(c *gin.Context) {
	json := QueueAddRequest{}

	c.Status(http.StatusBadRequest)

	if c.Bind(&json) == nil {
		var track *music.Track

		if json.TrackID != "" {
			track = music.GetTrackById(json.TrackID)
		} else if json.TrackName != "" {
			track = music.GetTrackByName(json.TrackName)
		}

		if track != nil {
			music.AddTrackToQueue(*track)
			c.Status(http.StatusOK)
		}
	}
}

func handleVolumeRequest(c *gin.Context) {
	json := VolumeSetRequest{}

	c.Status(http.StatusBadRequest)

	if c.Bind(&json) == nil {
		music.SetVolume(json.Volume)
		c.Status(http.StatusOK)
	}
}

func handleSkipRequest(c *gin.Context) {
	music.Skip()
	c.Status(http.StatusOK)
}
