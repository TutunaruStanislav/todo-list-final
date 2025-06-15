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

type UserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (input *UserInput) ToUser() *User {
	return &User{
		Username: input.Username,
		Password: input.Password,
	}
}
