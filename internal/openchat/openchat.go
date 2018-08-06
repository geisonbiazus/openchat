package openchat

type User struct {
	ID       string
	Username string
	Password string
	About    string
}

type UserRepository interface {
	Create(user User)
}

type IDGenerator interface {
	Generate() string
}
