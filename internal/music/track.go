package music

type Track struct {
	FileName string `json:"file_name" binding:"required"`
}
