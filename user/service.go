package user

import "golang.org/x/crypto/bcrypt"

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	GetUsers(userID int) ([]User, error)
	GetUserByID(input GetUserInput) (User, error)
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

func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email

	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (s *service) GetUsers(userID int) ([]User, error) {
	users, err := s.repository.FindAll()
	if err != nil {
		return users, err
	}

	return users, nil

}

func (s *service) GetUserByID(input GetUserInput) (User, error) {
	user, err := s.repository.FindByID(input.ID)

	if err != nil {
		return user, err
	}

	return user, nil
}
