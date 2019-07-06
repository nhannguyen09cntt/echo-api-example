package handler

import (
	"encoding/json"

	"github.com/labstack/echo"

	_ "forum-api/constant"
	models "forum-api/model"
)

type GroupHandler struct {
	BaseHandler
}

func NewGroupHandler() *GroupHandler {
	return &GroupHandler{}
}

func (h *GroupHandler) List() echo.HandlerFunc {
	return func(c echo.Context) error {
		groups, _, _ := models.NewGroups().List(25, 1)
		c.Response().Header().Set("Server", "4Rum")
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		return json.NewEncoder(c.Response()).Encode(groups)
	}
}

func (h *GroupHandler) Show() echo.HandlerFunc {
	return func(c echo.Context) error {
		identify := c.Param("identify")
		group, _ := models.NewGroups().One(identify)
		c.Response().Header().Set("Server", "4Rum")
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		return json.NewEncoder(c.Response()).Encode(group)
	}
}
