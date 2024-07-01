package main

import (
	"os"
	"todos/config"
	"todos/internal/controllers/todos"
	"todos/internal/controllers/users"
	"todos/internal/models"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	SecrateJWT := os.Getenv("SECRATE_JWT")
	e := echo.New() // initial echo
	setup := config.ImportSetting();
	connect, _ := config.ConnectDB(setup);

	connect.AutoMigrate(&models.Users{}, &models.Todos{})
	um := models.NewUserModels(connect);
	uc := users.NewUserController(um)

	tm := models.NewTodosModel(connect);
	tc := todos.NewTodosControllers(tm)

	e.POST("/register", uc.Register);
	e.POST("/login", uc.Login);

	// agar persingkat
	t := e.Group("/todos");
	t.Use(echojwt.WithConfig(
		echojwt.Config{
			SigningKey: []byte(SecrateJWT), // signing -> menggunakan key nya
			SigningMethod: jwt.SigningMethodHS256.Name,
		},
	))
	t.POST("", tc.CreateTodos);
	t.GET("", tc.GetTodos);
	t.PUT("/:id", tc.UpdateTodos);
	t.DELETE("/:id", tc.DeleteTodos);

	e.Pre(middleware.RemoveTrailingSlash()) // (wajib -> digunakan sebelum menthod nya di akses) => untuk menghapus slash yang berlebih di akhir endpoint
	e.Use(middleware.Logger()) // (wajib -> di gunakan setelah method nya di akses) digunakan untuk memunculkan console.log di terminal ketika sebuah end point di akses

	// CORS => di gunakan untuk membatasi arus komunikasi
	e.Use(middleware.CORSWithConfig(
		middleware.CORSConfig{
			// AllowOrigins: []string{"www.akbar.com"}, // hanya bisa di akses di "www.akbar.com" selain di url itu tidak bisa di akses
			// AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},	// method yang hanya bisa di akses
		},
	)) 

	e.Start(":8000")
}