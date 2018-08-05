package openchat

type User struct {
	Username string
	Password string
	About    string
}

type UserRepository interface {
	Create(user User)
}
