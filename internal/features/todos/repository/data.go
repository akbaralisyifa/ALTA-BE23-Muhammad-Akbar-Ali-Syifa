package repository

import (
	"todos/internal/features/todos"

	"gorm.io/gorm"
)

type Todos struct {
	gorm.Model
	UserID      uint   `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
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

func toTodoData(input todos.Todo) Todos {
	return Todos{
		UserID: input.UserID,
		Title: input.Title,
		Description: input.Description,
		Status: input.Status,
	}
}
