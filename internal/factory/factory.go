package factory

import (
	"todos/config"
	handlerTodo "todos/internal/features/todos/handler"
	repositoryTodo "todos/internal/features/todos/repository"
	"todos/internal/features/todos/services"
	"todos/internal/features/users/handler"
	"todos/internal/features/users/repository"
	"todos/internal/features/users/service"
	"todos/internal/routes"

	"github.com/labstack/echo/v4"
)

func InitalFactory(e *echo.Echo){
	setup := config.ImportSetting();
	connect, _ := config.ConnectDB(setup);

	connect.AutoMigrate(&repository.Users{}, &repositoryTodo.Todos{})
	um := repository.NewUserModel(connect);
	us := service.NewUserServices(um);
	uc := handler.NewUserController(us)

	tm := repositoryTodo.NewTodoModel(connect)
	ts := services.NewTodoSevices(tm)
	tc := handlerTodo.NewTodosControllers(ts)

	routes.InitialRoute(e, uc, tc);
}