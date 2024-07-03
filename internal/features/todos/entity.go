package todos

import "github.com/labstack/echo/v4"

type Todo struct {
	ID          uint
	UserID      uint
	Title       string
	Description string
	Status      bool
}

type Hendler interface {
	CreateTodos() echo.HandlerFunc
	GetTodos() 	echo.HandlerFunc
	UpdateTodos() echo.HandlerFunc
	DeleteTodos() echo.HandlerFunc
}

type Services interface {
	CreateTodos(newTodos Todo) error
	GetTodos(userID uint) (Todo, error)
	UpdateTodos(userID uint, id uint, newStatus bool)error
	DeleteTodos(userID uint, id uint) error

}

type Query interface {
	CreateTodos(newTodo Todo) error
	GetTodos(UserId uint) (Todo, error)
	UpdateTodos(UserID uint, id uint, newStatus bool) error
	DeleteTodos(UserID uint, id uint) error
}