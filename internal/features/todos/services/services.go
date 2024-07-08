package services

import "todos/internal/features/todos"

type TodoService struct {
	qry todos.Query
};

func NewTodoSevices(q todos.Query) todos.Services{
	return &TodoService{
		qry: q,
	}  
}

func (ts *TodoService) CreateTodos(newTodo todos.Todo)(error){
	err := ts.qry.CreateTodos(newTodo)

	if err != nil {
		return err;
	}

	return nil;
}

func (ts *TodoService) GetTodos(userID uint)( todos.Todo, error) {
	result, err := ts.qry.GetTodos(userID);

	if err != nil {
		return todos.Todo{}, err
	}

	return result, nil;
}

func (ts *TodoService) UpdateTodos(userID uint, id uint, newStatus bool)(error){

	err:= ts.qry.UpdateTodos(userID, id, newStatus);

	if err != nil {
		return err
	}

	return nil
}

func (ts *TodoService) DeleteTodos(userID uint, id uint) error {
	err := ts.qry.DeleteTodos(userID, id);

	if err != nil {
		return err
	}

	return nil;
}
