package entities

// auth service
// -register user-
type RequestRegisterUser struct {
	UserEmail    string
	UserName     string
	UserPassword string
}

// -Login user-
type RequestLoginUser struct {
	Email      string
	Password   string
	RememberMe bool
}

// -select user by id-
type RequestSelectSessionUserById struct {
	UserId string
}

type QuerySelectUserById struct {
	UserId string
}
