package data

import (
	"todolistapi/features/todo"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title           string
	IsActive        bool   `gorm:"default:1"`
	Priority        string `gorm:"default:'very-high'"`
	ActivityGroupID uint
}

func CoreToModel(data todo.Core) Todo {
	return Todo{
		Model: gorm.Model{
			ID:        data.ID,
			CreatedAt: data.CreatedAt,
			UpdatedAt: data.UpdatedAt,
		},
		Title:           data.Title,
		IsActive:        data.IsActive,
		Priority:        data.Priority,
		ActivityGroupID: data.ActivityGroupID,
	}
}

func ModelToCore(data Todo) todo.Core {
	return todo.Core{
		ID:              data.ID,
		Title:           data.Title,
		IsActive:        data.IsActive,
		Priority:        data.Priority,
		ActivityGroupID: data.ActivityGroupID,
		CreatedAt:       data.CreatedAt,
		UpdatedAt:       data.UpdatedAt,
	}
}
