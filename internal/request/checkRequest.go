package request

// CheckRequest check request struct
type CheckRequest struct {
	ID   int    `uri:"id" binding:"required"`
	Word string `uri:"word" binding:"required"`
}
