package models

import (
	"gorm.io/gorm"
)

type Todo struct {  // buat struct todo untuk tabel nya
	gorm.Model
	Activity string
	Mark	 bool
	Owner 	 uint
}

type TodoModel struct{ // lalu struct todo db nya
	db gorm.DB
}

func NewTodoModel(connection *gorm.DB) *TodoModel { // buat fungsi todo untuk terkoneksi ke DB
	return &TodoModel{
		db: *connection,
	}
};

func (tm *TodoModel) AddTodo(newData Todo)(Todo, error){
	newData.Mark = false;
	// buat ketika error
	err := tm.db.Create(&newData).Error;

	if err != nil {
		return Todo{}, err;
	}

	return newData, nil;
}

func (tm *TodoModel) UpdateTodo(ID uint, updatedData map[string]interface{})(Todo, error){
	var todo Todo;

	result := tm.db.Model(&todo).Where("ID = ?", ID).Updates(updatedData);
	
	if result.Error != nil {
		return Todo{}, result.Error
	}

	tm.db.First(&todo, ID)

	return todo, nil;
}

func (tm *TodoModel) DeleteTodo(ID uint)(Todo, error){
	
	var todo Todo;
	result := tm.db.Where("ID = ?", ID).Delete(&todo);

	if result.Error != nil {
		return Todo{}, result.Error;
	}

	return todo, nil;
}

func (tm *TodoModel)GetTodo()([]Todo, error){
	var todos []Todo;

	result := tm.db.Find(&todos);

	if result.Error != nil {
		return nil, result.Error
	}

	return todos, nil;

}