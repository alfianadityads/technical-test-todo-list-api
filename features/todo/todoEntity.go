package todo

import "time"

type Core struct {
	ID              uint
	Title           string
	Priority        string
	IsActive        bool
	ActivityGroupID uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
}