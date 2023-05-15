package activity

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID       uint
	Title    string
	Email    string
	CreateAt time.Time
	UpdateAt time.Time
}

type ActivityHandler interface {
	Create() echo.HandlerFunc
	GetOne() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}