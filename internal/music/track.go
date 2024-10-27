package music

import (
	"log"
	"os"
	"strings"

	"github.com/google/uuid"
)

type Track struct {
	ID       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	FileName string `json:"fileName" binding:"required"`
}

func GetTrackById(id string) *Track {
	for _, track := range Tracks {
		if track.ID == id {
			return &track
		}
	}

	return nil
}

func GetTrackByName(name string) *Track {
	for _, track := range Tracks {
		if track.Name == name {
			return &track
		}
	}

	return nil
}

func loadTracks() {
	files, err := os.ReadDir(tracksDirectory)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fileName := file.Name()
		trackName := strings.Split(fileName, ".")[0]

		track := Track{
			ID:       uuid.NewString(),
			Name:     trackName,
			FileName: fileName,
		}
		Tracks = append(Tracks, track)
	}

	log.Printf("Found %d track(s)", len(Tracks))
}
