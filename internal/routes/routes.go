package routes

import (
	"os"
	"todos/internal/features/todos"
	"todos/internal/features/users"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitialRoute(e *echo.Echo, uc users.Handler, tc todos.Hendler) {
	SecrateJWT := os.Getenv("SECRATE_JWT");
	
	e.POST("/register", uc.Register());
	e.POST("/login", uc.Login());

		// agar persingkat
		t := e.Group("/todos");
		t.Use(echojwt.WithConfig(
			echojwt.Config{
				SigningKey: []byte(SecrateJWT), // signing -> menggunakan key nya
				SigningMethod: jwt.SigningMethodHS256.Name,
			},
		))
		t.POST("", tc.CreateTodos());
		t.GET("", tc.GetTodos());
		t.PUT("/:id", tc.UpdateTodos());
		t.DELETE("/:id", tc.DeleteTodos());
	
		e.Pre(middleware.RemoveTrailingSlash()) // (wajib -> digunakan sebelum menthod nya di akses) => untuk menghapus slash yang berlebih di akhir endpoint
		e.Use(middleware.Logger()) // (wajib -> di gunakan setelah method nya di akses) digunakan untuk memunculkan console.log di terminal ketika sebuah end point di akses
	
		// CORS => di gunakan untuk membatasi arus komunikasi
		e.Use(middleware.CORSWithConfig(
			middleware.CORSConfig{
				// AllowOrigins: []string{"www.akbar.com"}, // hanya bisa di akses di "www.akbar.com" selain di url itu tidak bisa di akses
				// AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},	// method yang hanya bisa di akses
			},
		)) 
	
}