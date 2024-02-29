package refid

type RegisterRefIdInput struct {
	ID int `json:"id" binding:"required"`
}
