package controller

import (
	"fmt"
	"try/internal/models"
)

type UserController struct {
	model *models.UserModel
};

func NewUserController(md *models.UserModel) *UserController{
	return &UserController{
		model: md,
	}
}

func (uc *UserController) Login()(*models.User, error){
	var email, password string;
	fmt.Println("massukan Email");
	fmt.Scanln(&email);
	fmt.Println("Massukan Password");
	fmt.Scanln(&password);

	result, err := uc.model.Login(email,password);

	if err != nil {
		return &models.User{}, err;
	}

	return &result, nil;
}

func (uc *UserController) Register()(*models.User, error){
	var newData models.User;

	fmt.Print("Masukkan Nama :");
	fmt.Scanln(&newData.Name);
	fmt.Print("Masukkan Email :");
	fmt.Scanln(&newData.Email);
	fmt.Print("Masukkan Password :");
	fmt.Scanln(&newData.Password);
	fmt.Print("Masukkan No Hp :");
	fmt.Scanln(&newData.Phone);

	result, err := uc.model.Register(newData);

	if err != nil && !result {
		return &models.User{}, err;
	}

	return &newData, nil;
}