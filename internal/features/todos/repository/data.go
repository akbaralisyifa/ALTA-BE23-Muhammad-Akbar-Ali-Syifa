package repository

import (
	"todos/internal/features/todos"

	"gorm.io/gorm"
)

type Todos struct {
	gorm.Model
	UserID      uint   
	Title       string 
	Description string 
	Status      bool  
};

func (t *Todos) ToTodosEntity() todos.Todo{
	return todos.Todo{
		ID: 			t.ID,
		UserID: 		t.UserID,
		Title:			t.Title,
		Description:	t.Description,
		Status:			t.Status,
	}
}

func ToTodoData(input todos.Todo) Todos {
	return Todos{
		UserID: input.UserID,
		Title: input.Title,
		Description: input.Description,
		Status: input.Status,
	}
}
