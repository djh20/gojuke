package music

import (
	"log"
	"path/filepath"

	vlc "github.com/adrg/libvlc-go/v3"
)

var player *vlc.Player
var media *vlc.Media

func runPlayer() {
	var err error
	// Initialize libVLC. Additional command line arguments can be passed in
	// to libVLC by specifying them in the Init function.
	err = vlc.Init("--no-video", "--quiet")
	if err != nil {
		log.Fatal(err)
	}
	defer vlc.Release()

	// Create a new player.
	player, err = vlc.NewPlayer()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		player.Stop()
		player.Release()
	}()

	// Retrieve player event manager.
	manager, err := player.EventManager()
	if err != nil {
		log.Fatal(err)
	}

	// Register the media end reached event with the event manager.
	eventCallback := func(event vlc.Event, userData interface{}) {
		skipChannel <- PlaybackSkipEvent{}
	}

	eventID, err := manager.Attach(vlc.MediaPlayerEndReached, eventCallback, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer manager.Detach(eventID)

	for {
		select {
		case track := <-queueChannel:
			addTrackToQueue(track)

		case <-skipChannel:
			log.Println("Skipping track")
			playNextTrack()

		}
	}
}

func addTrackToQueue(track Track) {
	log.Printf("Appending %s to queue", track.FileName)
	Queue = append(Queue, track)

	if CurrentTrack == nil {
		playNextTrack()
	}
}

func playNextTrack() {
	var track Track

	player.Stop()

	if media != nil {
		media.Release()
		media = nil
	}

	CurrentTrack = nil

	if len(Queue) > 0 {
		track, Queue = Queue[0], Queue[1:]

		mediaPath := filepath.Join(tracksDirectory, track.FileName)

		var err error
		media, err = player.LoadMediaFromPath(mediaPath)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Playing", track.FileName)

		player.SetVolume(70)
		player.Play()

		CurrentTrack = &track
	}
}
