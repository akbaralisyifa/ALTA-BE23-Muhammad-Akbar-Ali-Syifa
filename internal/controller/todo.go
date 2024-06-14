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
	var todoID uint

	fmt.Println("UPDATE TODO");
	fmt.Println("Input '0' jika ingin batal");
	fmt.Println("Input ID kegiatan yang ingin di update !");
	fmt.Scanln(&todoID);

	if todoID != 0 {
		return true, nil
	}

	return false, nil;
}

func (tc *TodoController) DeleteTodo(id uint)(bool, error){
	var deleteData models.Todo;

	fmt.Println("input id aktivitas yang ingin di hapus");
	fmt.Scanln(&deleteData.ID);

	deleteData.Owner = id;
	_, err := tc.model.DeleteTodo(deleteData)

	if err != nil {
		return false, err;
	} 

	return true, nil
}	

func (tc *TodoController) FindTodos(ID uint) ([]map[string]any, error){
	
	var result []map[string]any;
	data, err := tc.model.FindTodo(ID);
	
	if err != nil {
		return nil, err
		}
		
	for _, todo := range data {
		todoMap := map[string]any{
			"id"		:todo.ID,
			"activity"	:todo.Activity,
			"owner"		:todo.Owner,
		}
		result = append(result, todoMap)
	}
	return result, nil
}

