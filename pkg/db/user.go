package db

// Base model for user
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

// Model for deserialization
type UserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// ToUser is function for converting a model into a basic user model
func (input *UserInput) ToUser() *User {
	return &User{
		Username: input.Username,
		Password: input.Password,
	}
}
