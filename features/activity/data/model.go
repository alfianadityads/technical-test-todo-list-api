package data

import "gorm.io/gorm"

type Activity struct {
	gorm.Model
	Title string
	Email string
	// Todo []data.Todo `gorm:"foreignkey:ActivityGroupID"`
}