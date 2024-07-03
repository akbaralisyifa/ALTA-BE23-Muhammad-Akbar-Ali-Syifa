package repository

import (
	"gorm.io/gorm"
)

type TodoModel struct {
	db gorm.DB
}

