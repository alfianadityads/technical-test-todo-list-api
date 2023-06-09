package route

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	activData "todolistapi/features/activity/data"
	activHandler "todolistapi/features/activity/handler"
	activService "todolistapi/features/activity/service"

	todoData "todolistapi/features/todo/data"
	todoHdl "todolistapi/features/todo/handler"
	todoSrv "todolistapi/features/todo/service"
)

func InitRoute(db *gorm.DB, e *echo.Echo) {
	activityData := activData.New(db)
	activityService := activService.New(activityData)
	activityHandler := activHandler.New(activityService)

	todoData := todoData.New(db)
	todoService := todoSrv.New(todoData)
	todoHandler := todoHdl.New(todoService)

	// Activity
	act := e.Group("/activity-groups")

	act.POST("", activityHandler.Create())
	act.GET("", activityHandler.GetAll())
	act.GET("/:id", activityHandler.GetOne())
	act.PATCH("/:id", activityHandler.Update())
	act.DELETE("/:id", activityHandler.Delete())

	// Todo
	todo := e.Group("/todo-items")

	todo.POST("", todoHandler.Create())
	todo.GET("/:id", todoHandler.GetOne())
	todo.GET("", todoHandler.GetAll())
	todo.PATCH("/:id", todoHandler.Update())
	todo.DELETE("/:id", todoHandler.Delete())
}
