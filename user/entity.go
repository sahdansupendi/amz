package user

import "time"

type User struct {
	ID             int
	Name           string
	NoHp           string
	Email          string
	Pwd            string
	AvatarFileName string
	Role           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (User) TableName() string {
	return "users"
}
