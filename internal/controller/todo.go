package controller

import (
	"fmt"
	"try/internal/models"
)

type TodoController struct {
	model *models.TodoModel
}

func NewTodoController(md *models.TodoModel) *TodoController{
	return &TodoController{
		model: md,
	}
}

func (tc *TodoController) AddTodo(id uint) (bool, error) {
	var newData models.Todo;

	fmt.Print("Add Todo :");
	fmt.Scanln(&newData.Activity);
	newData.Owner = id;

	_, err := tc.model.AddTodo(newData);

	if err != nil {
		return false, err;
	}
	return true, nil
};

func (tc *TodoController) UpdateTodo(ID uint) (bool, error){
	var updateData map[string]interface{};
	updateData = make(map[string]interface{});

	var mark bool
	fmt.Print("Set Mark :");
	fmt.Scanln(&mark);
	updateData["mark"] = mark;

	_, err := tc.model.UpdateTodo(ID, updateData);

	if err != nil {
		return false, err
	}

	return true, nil;
}

func (tc *TodoController) DeleteTodo(ID uint)(bool, error){
	
	fmt.Println("Are you sure to Delete ?");
	_, err := tc.model.DeleteTodo(ID);

	if err != nil{
		return false, err;
	}

	return true, nil
}	

func (tc *TodoController) GetTodo() ([]models.Todo, error){
	todos, err := tc.model.GetTodo();

	if err != nil {
		return nil, err;
	}

	return todos, nil
}

