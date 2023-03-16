package form

type Register struct {
	Email    string
	Password string
	Lang     string
}

type ConfirmRegister struct {
	Email    string
	Password string
	Code     string
}

type Login struct {
	Email    string
	Password string
}

type Token struct {
	Token string
}

type RecoveryPasswordRqt struct {
	Email string
}

type RecoveryPasswordConfirm struct {
	Email string
	Code  string
}

type ChangePassword struct {
	Email    string
	Password string
}
