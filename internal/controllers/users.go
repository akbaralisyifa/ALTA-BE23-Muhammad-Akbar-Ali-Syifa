package controllers

import (
	"todos/internal/helpers"
	"todos/internal/models"

	"github.com/labstack/echo/v4"
)

type UsersControllers struct {
	model *models.UserModels
}


func NewUserController(m *models.UserModels) *UsersControllers {
	return &UsersControllers{
		model: m,
	}
}


func(uc *UsersControllers) Register(c echo.Context) error {
	var input models.Users
	err := c.Bind(&input);

	if err != nil {
		return c.JSON(400, helpers.ResponseFormat(400, "input failed", nil));
	}

	// mengecek ketika input nya sudah di isi namun tidak tersimpan ke server
	_, err = uc.model.Register(input);
	if err != nil {
		return c.JSON(500, helpers.ResponseFormat(500, "server error", nil))
	}

	return c.JSON(201, helpers.ResponseFormat(201, "success register", nil));
}

func (uc *UsersControllers) Login(c echo.Context) error {
	var input models.Users;
	err := c.Bind(&input);

	if err != nil {
		return c.JSON(400, helpers.ResponseFormat(400, "input failed", nil));
	}

	result, err := uc.model.Login(input.Email, input.Password);

	if err != nil {
		return c.JSON(500, helpers.ResponseFormat(500, "Server Error", nil))
	}

	return c.JSON(201, helpers.ResponseFormat(201, "success login", result));
}