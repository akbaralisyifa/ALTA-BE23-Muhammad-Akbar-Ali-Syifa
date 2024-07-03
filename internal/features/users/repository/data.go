package repository

import (
	"todos/internal/features/todos/repository"
	"todos/internal/features/users"

	"gorm.io/gorm"
)

// untuk keperluan database
type Users struct {
	gorm.Model
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Phone    string  `json:"hp"`
	Todos    []repository.Todos `gorm:"foreignKey:UserID"`
}

func(u *Users) toUserEntity() users.User {
	return users.User{
		ID:			u.ID,	
		Username:	u.Username,
		Email:		u.Email,
		Password:	u.Password,
		Phone:		u.Phone,
	}
}

func toUserData(input users.User) Users {
	return Users{
		Username:	input.Username,
		Email:		input.Email,
		Password:	input.Password,
		Phone:		input.Phone,
	}
}
