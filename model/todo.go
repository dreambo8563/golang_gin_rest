package model

import (
	"github.com/jinzhu/gorm"
)

type (
	// TodoModel describes a todoModel type
	TodoModel struct {
		gorm.Model
		Title     string `form:"title" json:"title" binding:"required"`
		Completed int    `form:"completed" json:"completed" binding:"required"`
	}
	// TransformedTodo represents a formatted todo
	TransformedTodo struct {
		ID        uint   `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
)
