package main

import (
	"mini_project/config"
	"mini_project/route"

	"github.com/labstack/echo/v4"
)

func main() {
	db := config.InitDB()
	e := echo.New()

	route.NewRoute(e, db)

	e.Logger.Fatal(e.Start(":8080"))
}
