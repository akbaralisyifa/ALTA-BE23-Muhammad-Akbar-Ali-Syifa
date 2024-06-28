package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Username 	string `json:"username"`
	Email 		string	`json:"email"`
	Password 	string	`json:"password"`
	Phone		string	`json:"hp"`
	Todos 		Todos `gorm:"foreignKey:UserID"`
}

type UserModels struct {
	db *gorm.DB
}

func NewUserModels(connection *gorm.DB) *UserModels {
	return &UserModels{
		db : connection,
	}
}

func (um *UserModels) Register(newUser Users)(bool, error){
	err := um.db.Create(&newUser).Error;

	if err != nil {
		return false, err;
	};

	return true, nil;
}

func (um *UserModels) Login(email string, password string)(Users, error){
	var result Users;

	err := um.db.Where("email = ? AND password = ?", email, password).First(&result).Error;

	if err != nil {
		return Users{}, err
	}

	return result, nil;
}
