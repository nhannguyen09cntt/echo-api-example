package main

import (
	"forum-api/routers"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	routers.InitRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}
