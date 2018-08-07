package openchat

type User struct {
	ID       string
	Username string
	Password string
	About    string
}

var NoUser = User{}

type UserRepository interface {
	Create(user User)
	FindByUsername(username string) User
}

type IDGenerator interface {
	Generate() string
}

type Error struct {
	Field string
	Type  string
}

type ErrorsBuilder struct {
	Errors []Error
}

func NewErrorsBuilder() *ErrorsBuilder {
	return &ErrorsBuilder{}
}

func (e *ErrorsBuilder) Add(field, error_type string) {
	e.Errors = append(e.Errors, Error{Field: field, Type: error_type})
}
