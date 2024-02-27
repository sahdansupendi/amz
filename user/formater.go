package user

type UserFormatter struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	NoHp     string `json:"no_hp"`
	Email    string `json:"email"`
	Token    string `json:"token"`
	ImageUrl string `json:"image_url"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:       user.ID,
		Name:     user.Name,
		NoHp:     user.NoHp,
		Email:    user.Email,
		Token:    token,
		ImageUrl: user.AvatarFileName,
	}

	return formatter
}
