package form

type RegisterForm struct {
	Email    string
	Password string
}

type LoginForm struct {
	Email    string
	Password string
}

type TokenForm struct {
	Token string
}
