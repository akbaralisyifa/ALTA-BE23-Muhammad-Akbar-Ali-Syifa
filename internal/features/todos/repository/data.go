package repository

import "gorm.io/gorm"

type Todos struct {
	gorm.Model
	UserID      uint   `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
};


