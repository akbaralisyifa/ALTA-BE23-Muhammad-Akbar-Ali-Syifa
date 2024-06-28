package controllers

import (
	"strconv"
	"todos/internal/helpers"
	"todos/internal/models"

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
	var input models.Todos;
	err := c.Bind(&input);

	if err != nil {
		return c.JSON(400, helpers.ResponseFormat(400, "input failed", nil));
	}

	userID, err := strconv.Atoi(c.Param("UserID"))

	if err != nil {
		c.JSON(400, helpers.ResponseFormat(400, "Invalid User ID", nil));
	}

	err = tc.model.CreateTodos(uint(userID), input);

	if err != nil {
		return c.JSON(500, helpers.ResponseFormat(500, "error server", nil))
	 }

	return c.JSON(201, helpers.ResponseFormat(201, "success create todo", nil));
}

// GET TODOS
func(tc *TodosControllers) GetTodos(c echo.Context) error {
	
	userID, err := strconv.Atoi(c.Param("UserID"));

	if err != nil {
		return c.JSON(400, helpers.ResponseFormat(400, "user id failed", nil))
	}

	result, err := tc.model.GetTodos(uint(userID));

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
	var input models.Todos

	err := c.Bind(input.Status);

	if err != nil {
		return c.JSON(400, helpers.ResponseFormat(400, "input failed", nil))
	}

	userID, _ := strconv.Atoi(c.Param("UserID"));
	id, _ := strconv.Atoi(c.Param("id"));

	err = tc.model.UpdateTodos(uint(userID), uint(id), input.Status);

	if err != nil {
		return c.JSON(500, helpers.ResponseFormat(500, "server error", nil))
	}

	return c.JSON(200, helpers.ResponseFormat(200, "success update", nil))
}

// DELETE TODOS
func (tc *TodosControllers) DeleteTodos(c echo.Context) error {

	userID, err := strconv.Atoi(c.Param("UserID"));

	if err != nil {
		c.JSON(400, helpers.ResponseFormat(400, "user id failed", nil))
	}
	
	id, err := strconv.Atoi(c.Param("id"));
	
	if err != nil {
		c.JSON(400, helpers.ResponseFormat(400, "user id failed", nil))
	}


	err = tc.model.DeleteTodos(uint(userID), uint(id));

	if err != nil {
		return c.JSON(500, helpers.ResponseFormat(500, "server error", nil))
	}

	return c.JSON(200, helpers.ResponseFormat(200, "success delete", nil))
}