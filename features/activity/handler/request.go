package handler

import "todolistapi/features/activity"

type CreateActivityReq struct {
	Title string `validate:"required" json:"title"`
	Email string `validate:"required" json:"email"`
}

type UpdateActivityReq struct {
	Title string `validate:"required" json:"title"`
}

func ReqToCore(data interface{}) *activity.Core {
	res := activity.Core{}

	switch data.(type) {
	case CreateActivityReq:
		cnv := data.(CreateActivityReq)
		res.Title = cnv.Title
		res.Email = cnv.Email
	case UpdateActivityReq:
		cnv := data.(UpdateActivityReq)
		res.Title = cnv.Title
	default:
		return nil
	}

	return &res
}