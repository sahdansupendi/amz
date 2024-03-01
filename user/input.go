package user

type RegisterUserInput struct {
	Name  string `json:"name" binding:"required"`
	NoHp  string `json:"no_hp"`
	Email string `json:"email" binding:"required,email"`
	Pwd   string `json:"password" binding:"required"`
}

type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}

type GetUserInput struct {
	ID int `uri:"id" binding:"required"`
}
