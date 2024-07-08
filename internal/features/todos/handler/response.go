package handler

import "todos/internal/features/todos"

type ResponseTodo struct {
	UserID      uint   `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

func ToResponseGetTodo(input todos.Todo) ResponseTodo {
	return ResponseTodo{
		Title: input.Title,
		Description: input.Description,
		Status: input.Status,
	}
}