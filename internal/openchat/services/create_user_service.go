package services

import "github.com/geisonbiazus/openchat/internal/openchat"

type CreateUserInput struct {
	Username string
	Password string
	About    string
}

type CreateUserService struct {
	UserRepository openchat.UserRepository
}

func NewCreateUserService(userRepository openchat.UserRepository) *CreateUserService {
	return &CreateUserService{
		UserRepository: userRepository,
	}
}

func (s *CreateUserService) Run(input CreateUserInput) {
	s.UserRepository.Create(s.newUser(input))
}

func (s *CreateUserService) newUser(input CreateUserInput) openchat.User {
	return openchat.User{
		Username: input.Username,
		Password: input.Password,
		About:    input.About,
	}
}
