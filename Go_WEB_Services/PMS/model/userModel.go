package model

type UserRegistration struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Mobile   int    `json:"mobile"`
	Password string `json:"password"`
}
