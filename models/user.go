package models

type User struct {
	ID           int64  `json:"-"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	PasswordHash string `json:"-"`
	Role         int64  `json:"role"`
	PaypalEmail  string `json:"paypalEmail"`
}

type Authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	Role        int64  `json:"role"`
	Email       string `json:"email"`
	TokenString string `json:"token"`
}
