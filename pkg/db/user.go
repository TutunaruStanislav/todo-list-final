package db

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewUser() *User {
	return &User{
		Username: "User-1",
		Password: "",
	}
}
