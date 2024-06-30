package todos

import (
	"strconv"
	"todos/internal/helpers"
	"todos/internal/models"
	"todos/internal/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type TodosControllers struct {
	model *models.TodosModel
}

func NewTodosControllers(m *models.TodosModel) *TodosControllers {
	return &TodosControllers{
		model: m,
	}
}


// CREATE TODOS
func(tc *TodosControllers) CreateTodos(c echo.Context) error {

	// mendapatkan user ID dari token jwt
	var userID = utils.DecodeToken(c.Get("user").(*jwt.Token))

	// requeste todo
	var input TodosRequest;
	err := c.Bind(&input);

	if err != nil {
		return c.JSON(400, helpers.ResponseFormat(400, "input failed", nil));
	}

	err = tc.model.CreateTodos(ToRequestModelTodo(input, userID));

	if err != nil {
		return c.JSON(500, helpers.ResponseFormat(500, "error server", nil))
	}

	return c.JSON(201, helpers.ResponseFormat(201, "success create todo", nil));
}

// GET TODOS
func(tc *TodosControllers) GetTodos(c echo.Context) error {
	// get id from token
	var userID = utils.DecodeToken(c.Get("user").(*jwt.Token))

	result, err := tc.model.GetTodos(userID);

	if len(result) == 0 {
		return c.JSON(400, helpers.ResponseFormat(400, "data empty", nil))
	}

	if err != nil {
		return c.JSON(500, helpers.ResponseFormat(500, "server error", nil))
	}
	
	return c.JSON(200, helpers.ResponseFormat(200, "success", result))
}

// BAGIAN UPDATE TODOS
func(tc *TodosControllers) UpdateTodos(c echo.Context) error {
	// get id from token
	var userID = utils.DecodeToken(c.Get("user").(*jwt.Token))

	var input TodoUpdateRequeste;
	err := c.Bind(&input);

	if err != nil {
		return c.JSON(400, helpers.ResponseFormat(400, "input failed", nil))
	}

	// mendapatkan parameter dari url
	paramId := c.Param("id");
	id, _ := strconv.Atoi(paramId)

	err = tc.model.UpdateTodos(userID, uint(id), input.Status);

	if err != nil {
		return c.JSON(500, helpers.ResponseFormat(500, "server error", nil))
	}

	return c.JSON(200, helpers.ResponseFormat(200, "success update", nil))
}

// DELETE TODOS
func (tc *TodosControllers) DeleteTodos(c echo.Context) error {

	// get id from token
	var userID = utils.DecodeToken(c.Get("user").(*jwt.Token))
	// get param
	ParamId := c.Param("id");
	id, _ := strconv.Atoi(ParamId)

	err := tc.model.DeleteTodos(userID, uint(id));

	if err != nil {
		return c.JSON(500, helpers.ResponseFormat(500, "server error", nil))
	}

	return c.JSON(200, helpers.ResponseFormat(200, "success delete", nil))
}