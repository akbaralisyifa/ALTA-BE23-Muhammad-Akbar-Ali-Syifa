package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Email string
	Password string
	Phone string
	BirthDate time.Time `gorm:"type:date"`
}

type UserModel struct {
	db *gorm.DB
};

func NewUserModel(connection *gorm.DB) *UserModel{
	return &UserModel{
		db: connection,
	}
}

func (userModel *UserModel) Login(email, password string) (User, error){
	var result User;
	// ketika error
	err := userModel.db.Where("email = ? AND password = ?", email, password).First(&result).Error;
	if err != nil {
		return User{}, err;
	}

	return result, nil;
};


func (userModel *UserModel) Register(newUser User) (bool, error){
	err:= userModel.db.Create(&newUser).Error;
	if err != nil {
		return false, err;
	}

	return true, nil
}