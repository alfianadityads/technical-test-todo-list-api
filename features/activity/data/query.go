package data

import (
	"errors"
	"fmt"
	"log"
	"todolistapi/features/activity"

	"gorm.io/gorm"
)

type activityQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) activity.ActivityData {
	return &activityQuery{
		db: db,
	}
}

// Create implements activity.ActivityData
func (aq *activityQuery) Create(newActivity activity.Core) (activity.Core, error) {
	cnv := CoreToModel(newActivity)
	err := aq.db.Create(&cnv).Error
	if err != nil {
		log.Println("Query create a new activity error : ", err.Error())
		return newActivity, err
	}
	newActivity.ID = cnv.ID

	return newActivity, nil
}

// Delete implements activity.ActivityData
func (aq *activityQuery) Delete(activityID uint) error {
	qryDelete := aq.db.Delete(&Activity{}, activityID)

	affRow := qryDelete.RowsAffected

	if affRow <= 0 {
		log.Println("No rows affected")
		msg := fmt.Sprintf("Activity with ID %d Not Found", activityID)
		return errors.New(msg)
	}

	return nil
}

// GerAll implements activity.ActivityData
func (aq *activityQuery) GetAll() ([]activity.Core, error) {
	allActivity := []Activity{}

	err := aq.db.Find(&allActivity).Error
	if err != nil {
		log.Println("Query get All activities error : ", err.Error())
		return []activity.Core{}, err
	}
	result := []activity.Core{}
	for _, val := range allActivity {
		result = append(result, ModelToCore(val))
	}

	return result, nil
}

// GetOne implements activity.ActivityData
func (aq *activityQuery) GetOne(activityID uint) (activity.Core, error) {
	activ := Activity{}

	err := aq.db.Where("id = ?", activityID).First(&activ).Error
	if err != nil {
		log.Println("Query get activity by ID error : ", err.Error())
		return activity.Core{}, err
	}

	return ModelToCore(activ), nil
}

// Update implements activity.ActivityData
func (aq *activityQuery) Update(activityID uint, updateActivity activity.Core) (activity.Core, error) {
	cnvUpdated := CoreToModel(updateActivity)
	qry := aq.db.Model(Activity{}).Where("id = ?", activityID).Updates(&cnvUpdated)
	err := qry.Error

	affRow := qry.RowsAffected

	if affRow <= 0 {
		log.Println("No rows affected")
		msg := fmt.Sprintf("Activity with ID %d Not Found", activityID)
		return activity.Core{}, errors.New(msg)
	}

	if err != nil {
		log.Println("Query update activity by ID error : ", err.Error())
		return activity.Core{}, errors.New("Error")
	}

	var updatedRow Activity
	aq.db.First(&updatedRow, "id = ?", activityID)

	return ModelToCore(updatedRow), nil
}
