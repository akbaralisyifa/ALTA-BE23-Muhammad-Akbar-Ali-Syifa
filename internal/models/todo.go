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

func (tm *TodoModel) UpdateTodo(userID, todoID uint)(error){
	query := tm.db.Model(&Todo{}).Where("ID = ? AND owner = ?", userID, todoID).Update("mark", true);
	return query.Error;
}

func (tm *TodoModel) DeleteTodo(deleteTodo Todo)(Todo, error){
	
	// query := ` UPDATE "be24"."todos" SET "deleted_at"= ?
	// WHERE (owner = ? AND activity = ?) AND "todos"."deleted_at" IS NULL `
	// err := tm.db.Exec(query, &deleteData.UpdatedAt, &deleteData.Owner, &deleteData.Activity).Error
	query := tm.db.Delete(&deleteTodo);

	// mengecek klo ada error
	if query.Error != nil {
		return Todo{}, query.Error;
	}
	// mengecek kalo id yang di hapus di database nya itu tidak ada
	if query.RowsAffected < 1 {
		return Todo{}, gorm.ErrRecordNotFound;
	}

	return deleteTodo, nil;
}

func (tm *TodoModel)FindTodo(owner uint)([]Todo, error){
	var todos []Todo;

	query := tm.db.Where("owner = ?", owner).Find(&todos);

	if query.Error != nil {
		return nil, query.Error
	}

	return todos, nil 

}