package user

import "golang.org/x/crypto/bcrypt"

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.NoHp = input.NoHp
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Pwd), bcrypt.MinCost)

	if err != nil {
		return user, err
	}

	user.Pwd = string(passwordHash)
	user.Role = "user"

	newUser, err := s.repository.Save(user)

	if err != nil {
		return user, nil
	}

	return newUser, nil
}