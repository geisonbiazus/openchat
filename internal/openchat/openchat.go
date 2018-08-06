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
