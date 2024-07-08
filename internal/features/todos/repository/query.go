package repository

import (
	"todos/internal/features/todos"

	"gorm.io/gorm"
)

type TodoModel struct {
	db gorm.DB
}

func NewTodoModel(connection *gorm.DB) todos.Query{
	return &TodoModel{
		db: *connection,
	}
}

func (tm *TodoModel) CreateTodos(newTodo todos.Todo)(error){
	cvrData := ToTodoData(newTodo)
	err := tm.db.Create(&cvrData).Error;
	return err;
}

func (tm *TodoModel) GetTodos(UserID uint)(todos.Todo, error){
	var result Todos;
	err := tm.db.Where("user_id = ?", UserID).Find(&result).Error;

	if err != nil {
		return todos.Todo{}, err;
	}

	return result.ToTodosEntity(), nil;
}

func (tm *TodoModel) UpdateTodos(userID uint, id uint, newStatus bool)(error){

	qry := tm.db.Model(Todos{}).Where("user_id = ? AND id = ?", userID, id).Update("status", newStatus)

	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return nil;
};

func (tm *TodoModel) DeleteTodos(userID uint, id uint)(error){
	qry := tm.db.Where("user_id = ? AND id = ?", userID, id).Delete(Todos{});

	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return nil
}