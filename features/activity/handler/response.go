package handler

import (
	"time"
	"todolistapi/features/activity"
)

type ActivityResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CoreToResp(data activity.Core) ActivityResponse {
	return ActivityResponse{
		ID:        data.ID,
		Title:     data.Title,
		Email:     data.Email,
		CreatedAt: data.CreateAt,
		UpdatedAt: data.UpdateAt,
	}
}

func CoreToRespArr(data []activity.Core) []ActivityResponse {
	res := []ActivityResponse{}
	for _, val := range data {
		tmp := CoreToResp(val)
		res = append(res, tmp)
	}

	return res
}
