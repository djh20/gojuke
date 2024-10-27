package music

import (
	"github.com/djh20/openjukebox/internal/storage"
)

var CurrentTrack *Track
var Queue []Track

// var tracks []Track
var tracksDirectory string

// var queue = make(chan Track, 50)
var queueChannel = make(chan Track, 10)
var skipChannel = make(chan PlaybackSkipEvent, 10)

func Init() {
	tracksDirectory = storage.GetDataSubDirectory("tracks")
}

func Run() {
	runPlayer()
}

func Skip() {
	skipChannel <- PlaybackSkipEvent{}
}

func AddTrackToQueue(track Track) {
	queueChannel <- track
}
