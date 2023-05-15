package service

import (
	"todolistapi/features/activity"
	"todolistapi/helper"

	"github.com/go-playground/validator"
)

type activityService struct {
	qry activity.ActivityData
	vld *validator.Validate
}

func New(ad activity.ActivityData) activity.ActivityService {
	return &activityService{
		qry: ad,
		vld: validator.New(),
	}
}

// Create implements activity.ActivityService
func (as *activityService) Create(newActivity activity.Core) (activity.Core, error) {
	err := helper.Validation(newActivity)
	if err != nil {
		return newActivity, err
	}

	res, err := as.qry.Create(newActivity)
	if err != nil {
		return res, err
	}

	return res, nil
}

// Delete implements activity.ActivityService
func (as *activityService) Delete(activityID uint) error {
	err := as.qry.Delete(activityID)
	if err != nil {
		return err
	}

	return nil
}

// GetAll implements activity.ActivityService
func (as *activityService) GetAll() ([]activity.Core, error) {
	panic("unimplemented")
}

// GetOne implements activity.ActivityService
func (as *activityService) GetOne(activityID uint) (activity.Core, error) {
	panic("unimplemented")
}

// Update implements activity.ActivityService
func (as *activityService) Update(activityID uint, updateActivity activity.Core) (activity.Core, error) {
	panic("unimplemented")
}
