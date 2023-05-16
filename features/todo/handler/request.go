package handler

type TodoRequest struct {
	Title           string `validate:"required" json:"title"`
	IsActive        bool   `json:"is_active"`
	ActivityGroupID uint   `validate:"required" json:"activity_group_id"`
}

type TodoUpdateRequest struct {
	Title    string `json:"title"`
	Priority string `json:"priority"`
	IsActive bool   `json:"is_active"`
}

