package entities

type ResponRegisterUser struct {
	UserName  string
	UserEmail string
}

type ResponLoginUser struct {
	UserId  string
	Session string
}

type ResponSelectSessionUserById struct {
	UserSession string
	RememberMe  bool
}
