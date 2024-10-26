package web

type queueAddRequest struct {
	Track string `json:"track" binding:"required"`
}
