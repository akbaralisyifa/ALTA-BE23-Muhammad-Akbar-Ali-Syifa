package handler

import (
	"todos/internal/features/todos"
)

// membuat requeste yang di perlukan
type TodosRequest struct{
	Title 		string	`json:"title"`
	Description string	`json:"description"`
	Status 		bool	`json:"status"`
}

type TodoUpdateRequeste struct {
	Status 	bool `json:"status"`
}

// fungsi request nya
func ToRequestModelTodo(tr TodosRequest, userID uint) todos.Todo {
	return todos.Todo{
		UserID: userID,
		Title: tr.Title,
		Description: tr.Description,
		Status: tr.Status,
	}
}