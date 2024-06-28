package models

import "gorm.io/gorm"

type Todos struct {
	gorm.Model
	UserID 		uint	
	Title 		string	`json:"title"`
	Description string	`json:"description"`
	Status 		bool	`json:"status"`
}

type TodosModel struct {
	db *gorm.DB
}

func NewTodosModel(connection *gorm.DB) *TodosModel {
	return &TodosModel{
		db : connection,
	}
}


func (tm *TodosModel) CreateTodos(userID uint, newTodo Todos)error{

	err := tm.db.Where("UserID = ?", userID).Create(&newTodo).Error;

	if err != nil {
		return err
	};

	return	nil;
}

func (tm *TodosModel) GetTodos(userID uint)([]Todos, error){
	var todos []Todos;

	err := tm.db.Where("UserID = ?", userID).Find(&todos).Error;

	if err != nil {
		return []Todos{}, err
	};

	return todos, nil;
}

func (tm *TodosModel) UpdateTodos( userID uint, id uint, newStatus any)(error){
	var todos Todos
	err := tm.db.Model(&todos).Where("UserID = ? AND id = ?", userID, id).Update("status", newStatus).Error;

	if err != nil {
		return err
	}

	return nil;
}

func(tm *TodosModel) DeleteTodos(userID uint, id uint) error {
	query := tm.db.Where("UserID = ? AND id = ?", userID, id).Delete(id);

	if query.RowsAffected < 1 {
		return query.Error
	}

	if query.Error != nil {
		return query.Error
	}

	return nil
}