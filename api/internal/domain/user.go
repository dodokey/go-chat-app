package domain

type User struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type UserRepository interface {
    Save(user *User) error
    FindByUsername(username string) (*User, error)
}