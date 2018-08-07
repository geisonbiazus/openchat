package services

import "github.com/geisonbiazus/openchat/internal/openchat"

const (
	StatusSuccess = "SUCCESS"
	StatusError   = "ERROR"
)

type CreateUserInput struct {
	Username string
	Password string
	About    string
}

type CreateUserOutput struct {
	Status   string
	ID       string
	Username string
	About    string
	Errors   []openchat.Error
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

func (s *CreateUserService) Run(input CreateUserInput) CreateUserOutput {
	user := s.newUser(input)

	if errors := s.validateUser(user); len(errors) > 0 {
		return s.buildInvalidOutput(errors)
	}

	s.UserRepository.Create(user)
	return s.buildSuccessOutput(user)
}

func (s *CreateUserService) newUser(input CreateUserInput) openchat.User {
	return openchat.User{
		ID:       s.IDGenerator.Generate(),
		Username: input.Username,
		Password: input.Password,
		About:    input.About,
	}
}

func (s *CreateUserService) validateUser(user openchat.User) []openchat.Error {
	builder := openchat.NewErrorsBuilder()

	if user.Username == "" {
		builder.Add("username", "REQUIRED")
	}

	if user.Password == "" {
		builder.Add("password", "REQUIRED")
	}

	if s.usernameIsTaken(user.Username) {
		builder.Add("username", "ALREADY_TAKEN")
	}

	return builder.Errors
}

func (s *CreateUserService) usernameIsTaken(username string) bool {
	user := s.UserRepository.FindByUsername(username)
	return user != openchat.NoUser
}

func (s *CreateUserService) buildInvalidOutput(errors []openchat.Error) CreateUserOutput {
	return CreateUserOutput{
		Status: StatusError,
		Errors: errors,
	}
}

func (s *CreateUserService) buildSuccessOutput(user openchat.User) CreateUserOutput {
	return CreateUserOutput{
		Status:   StatusSuccess,
		ID:       user.ID,
		Username: user.Username,
		About:    user.About,
	}
}
