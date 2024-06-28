package main

import (
	"todos/config"
	"todos/internal/controllers"
	"todos/internal/models"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New() // initial echo
	setup := config.ImportSetting();
	connect, _ := config.ConnectDB(setup);

	// connect.AutoMigrate(&models.Users{}, &models.Todos{})
	um := models.NewUserModels(connect);
	uc := controllers.NewUserController(um)

	tm := models.NewTodosModel(connect);
	tc := controllers.NewTodosControllers(tm)

	e.POST("/register", uc.Register);
	e.POST("/login", uc.Login);
	e.POST("/todos/:userID", tc.CreateTodos);
	e.GET("/todos/:userID", tc.GetTodos);
	e.PUT("/todos/:userID/:id", tc.UpdateTodos);
	e.DELETE("/todos/:userID/:id", tc.DeleteTodos);
	
}