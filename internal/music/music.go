package music

import "github.com/djh20/openjukebox/internal/storage"

var Tracks []Track
var CurrentTrack *Track
var Queue []Track

// var queue = make(chan Track, 50)
var queueChannel = make(chan Track, 10)
var skipChannel = make(chan struct{}, 10)
var volumeChannel = make(chan int, 10)

var tracksDirectory string

func Init() {
	tracksDirectory = storage.GetDataSubDirectory("tracks")

	loadTracks()
}

func Run() {
	runPlayer()
}

func SetVolume(volume int) {
	volumeChannel <- volume
}

func Skip() {
	skipChannel <- struct{}{}
}

func AddTrackToQueue(track Track) {
	queueChannel <- track
}
