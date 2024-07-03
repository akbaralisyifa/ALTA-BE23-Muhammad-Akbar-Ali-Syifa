package handler

import (
	"todos/internal/features/users"
	"todos/internal/helpers"

	"github.com/labstack/echo/v4"
)

type UsersControllers struct {
	srv users.Services
}

func NewUserController(s users.Services) users.Handler {
	return &UsersControllers{
		srv: s,
	}
}

func(uc *UsersControllers) Register() echo.HandlerFunc{

	return func(c echo.Context) error {
		var input RegisterRequest
		err := c.Bind(&input);
	
		if err != nil {
			c.Logger().Error("Register Error", err.Error())
			return c.JSON(400, helpers.ResponseFormat(400, "input failed", nil));
		}
	
		err = uc.srv.Register(ToModelUsers(input))
		
		if err != nil {
			return c.JSON(500, helpers.ResponseFormat(500, "server error", nil))
		}
	
		return c.JSON(201, helpers.ResponseFormat(201, "success register", nil));
	}
}


func (uc *UsersControllers) Login() echo.HandlerFunc{
	return func(c echo.Context) error {
		// ketika login -> gunakan pake struct login request yang di buat dari pada gunakan model users;
		// untuk menghemat tabel 
		var input LoginRequest;
		err := c.Bind(&input);

		if err != nil {
			return c.JSON(400, helpers.ResponseFormat(400, "input failed", nil));
		}

		result, token, err := uc.srv.Login(input.Email, input.Password)

		if err != nil {
			return c.JSON(500, helpers.ResponseFormat(500, "server error", nil));
		}

		return c.JSON(201, helpers.ResponseFormat(201, "success login", ToLoginResponse(result, token)));
	}
}