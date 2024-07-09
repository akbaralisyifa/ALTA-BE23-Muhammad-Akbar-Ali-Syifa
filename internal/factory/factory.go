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
	"todos/internal/utils"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func InitalFactory(e *echo.Echo){
	setup := config.ImportSetting();
	connect, _ := config.ConnectDB(setup);

	connect.AutoMigrate(&repository.Users{}, &repositoryTodo.Todos{})
	vldt := utils.NewValidatorUtility(*validator.New())
	jwt := utils.NewJwtUtility();
	pw := utils.NewGenertePassword()
	um := repository.NewUserModel(connect);
	us := service.NewUserServices(um, vldt, jwt, pw);
	uc := handler.NewUserController(us)

	dcJwt := utils.NewJwtUtility()
	tm := repositoryTodo.NewTodoModel(connect)
	ts := services.NewTodoSevices(tm)
	tc := handlerTodo.NewTodosControllers(ts, dcJwt)

	routes.InitialRoute(e, uc, tc);
}