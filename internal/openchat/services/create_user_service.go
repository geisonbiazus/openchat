package services

import "github.com/geisonbiazus/openchat/internal/openchat"

type CreateUserInput struct {
	Username string
	Password string
	About    string
}

type CreateUserService struct {
	UserRepository openchat.UserRepository
	IDGenerator    openchat.IDGenerator
}

func NewCreateUserService(
	idGen openchat.IDGenerator, userRep openchat.UserRepository,
) *CreateUserService {
	return &CreateUserService{
		IDGenerator:    idGen,
		UserRepository: userRep,
	}
}

func (s *CreateUserService) Run(input CreateUserInput) {
	s.UserRepository.Create(s.newUser(input))
}

func (s *CreateUserService) newUser(input CreateUserInput) openchat.User {
	return openchat.User{
		ID:       s.IDGenerator.Generate(),
		Username: input.Username,
		Password: input.Password,
		About:    input.About,
	}
}
