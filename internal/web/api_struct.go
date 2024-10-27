package web

type QueueAddRequest struct {
	TrackID   string `json:"trackId"`
	TrackName string `json:"trackName"`
}

type VolumeSetRequest struct {
	Volume int `json:"volume" binding:"required"`
}
