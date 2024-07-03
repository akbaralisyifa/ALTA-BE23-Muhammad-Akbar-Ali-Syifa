package handler

import (
	"todos/internal/features/users"
)

// struct untuk penginputan yang perlu saja (yang perlu di isi)
// lalu ganti controller login nya menjadi truct ini bukan struct dari users
type LoginRequest struct{
	Email 		string	`json:"email"`
	Password 	string	`json:"password"`
}
type RegisterRequest struct{
	Username 	string `json:"username"`
	Email 		string	`json:"email"`
	Password 	string	`json:"password"`
	Phone		string	`json:"hp"`
	// Todos 		[]Todos `gorm:"foreignKey:UserID"`
}

func ToModelUsers(r RegisterRequest) users.User{
	return users.User{
		Username: 	r.Username,
		Email: 		r.Email,
		Password: 	r.Password,
		Phone: 		r.Phone,
	}
}