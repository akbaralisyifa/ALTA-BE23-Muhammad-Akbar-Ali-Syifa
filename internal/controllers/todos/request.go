package todos

import "todos/internal/models"

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
func ToRequestModelTodo(tr TodosRequest, userID uint) models.Todos {
	return models.Todos{
		UserID: userID,
		Title: tr.Title,
		Description: tr.Description,
		Status: tr.Status,
	}
}