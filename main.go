package main

import (
	"log"
	"todolistapi/app/config"
	"todolistapi/app/database"
	"todolistapi/app/route"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.InitConfig()
	db := database.InitDB(*cfg)

	database.Migrate(db)

	e := echo.New()

	route.InitRoute(db, e)

	if err := e.Start(":3030"); err != nil {
		log.Println(err.Error())
	}
}
