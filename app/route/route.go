package route

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	activData "todolistapi/features/activity/data"
	activHandler "todolistapi/features/activity/handler"
	activService "todolistapi/features/activity/service"
)

func InitRoute(db *gorm.DB, e *echo.Echo) {
	activityData := activData.New(db)
	activityService := activService.New(activityData)
	activityHandler := activHandler.New(activityService)

	// Activity
	act := e.Group("/activity-groups")

	act.POST("", activityHandler.Create())
	act.GET("", activityHandler.GetAll())
	act.GET("/:id", activityHandler.GetOne())
	act.PATCH("/:id", activityHandler.Update())
	act.DELETE("/:id", activityHandler.Delete())
}
