package producer

type Job struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Created     int64  `json:"created"`
}
