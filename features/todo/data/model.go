package data

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title           string
	IsActive        bool   `gorm:"default:1"`
	Priority        string `gorm:"default:'very-high'"`
	ActivityGroupID uint
}
