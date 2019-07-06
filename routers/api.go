package routers

import (
	handler "forum-api/handler"

	"github.com/labstack/echo"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/group", handler.NewGroupHandler().List())
	e.GET("/group/:identify", handler.NewGroupHandler().Show())
}
