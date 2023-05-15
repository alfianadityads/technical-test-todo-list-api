package data

import (
	"todolistapi/features/activity"

	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model
	Title string
	Email string
	// Todo []data.Todo `gorm:"foreignkey:ActivityGroupID"`
}

func CoreToModel(data activity.Core) Activity {
	return Activity{
		Model: gorm.Model{
			ID:        data.ID,
			CreatedAt: data.CreateAt,
			UpdatedAt: data.UpdateAt,
		},
		Title: data.Title,
		Email: data.Email,
	}
}

func ModelToCore(data Activity) activity.Core {
	return activity.Core{
		ID:       data.ID,
		Title:    data.Title,
		Email:    data.Email,
		CreateAt: data.CreatedAt,
		UpdateAt: data.UpdatedAt,
	}
}
