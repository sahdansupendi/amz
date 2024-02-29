package user

type UserFormatter struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	NoHp     string `json:"no_hp"`
	Email    string `json:"email"`
	Token    string `json:"token"`
	ImageUrl string `json:"image_url"`
}

type ListUserFormatter struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	NoHp  string `json:"no_hp"`
	Email string `json:"email"`
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

func FormatListUsers(user User) ListUserFormatter {
	UserFormatter := ListUserFormatter{}
	UserFormatter.ID = user.ID
	UserFormatter.Name = user.Name
	UserFormatter.NoHp = user.NoHp
	UserFormatter.Email = user.Email
	return UserFormatter
}

func FormatUsers(users []User) []ListUserFormatter {
	usersFormatter := []ListUserFormatter{}

	for _, user := range users {
		userFormatter := FormatListUsers(user)
		usersFormatter = append(usersFormatter, userFormatter)
	}

	return usersFormatter
}
