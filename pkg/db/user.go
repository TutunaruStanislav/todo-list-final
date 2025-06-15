package db

type User struct {
	Username string `json:"username"`
	Password string `json:"-"`
}

func NewUser() *User {
	return &User{
		Username: "User-1",
		Password: "",
	}
}
