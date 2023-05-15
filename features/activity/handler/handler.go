package handler

import (
	"todolistapi/features/activity"

	"github.com/labstack/echo/v4"
)

type activHandler struct {
	srv activity.ActivityService
}

func New(srv activity.ActivityService) activity.ActivityHandler {
	return &activHandler{
		srv: srv,
	}
}

// Create implements activity.ActivityHandler
func (ah *activHandler) Create() echo.HandlerFunc {
	panic("unimplemented")
}

// Delete implements activity.ActivityHandler
func (ah *activHandler) Delete() echo.HandlerFunc {
	panic("unimplemented")
}

// GetAll implements activity.ActivityHandler
func (ah *activHandler) GetAll() echo.HandlerFunc {
	panic("unimplemented")
}

// GetOne implements activity.ActivityHandler
func (ah *activHandler) GetOne() echo.HandlerFunc {
	panic("unimplemented")
}

// Update implements activity.ActivityHandler
func (ah *activHandler) Update() echo.HandlerFunc {
	panic("unimplemented")
}
