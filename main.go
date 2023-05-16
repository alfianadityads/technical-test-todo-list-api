package main

import (
	"todolistapi/app/config"
	"todolistapi/app/database"

	// "github.com/labstack/echo/v4"
)

func main() {
	cfg := config.InitConfig()
	db := database.InitDB(*cfg)

	database.Migrate(db)

	// e := echo.New()

	// route
}