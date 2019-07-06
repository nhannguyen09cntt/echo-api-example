<a href="https://echo.labstack.com"><img height="80" src="https://cdn.labstack.com/images/echo-logo.svg"></a>

[![Sourcegraph](https://sourcegraph.com/github.com/labstack/echo/-/badge.svg?style=flat-square)](https://sourcegraph.com/github.com/labstack/echo?badge)
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/labstack/echo)
[![Go Report Card](https://goreportcard.com/badge/github.com/labstack/echo?style=flat-square)](https://goreportcard.com/report/github.com/labstack/echo)
[![Build Status](http://img.shields.io/travis/labstack/echo.svg?style=flat-square)](https://travis-ci.org/labstack/echo)
[![Codecov](https://img.shields.io/codecov/c/github/labstack/echo.svg?style=flat-square)](https://codecov.io/gh/labstack/echo)
[![Join the chat at https://gitter.im/labstack/echo](https://img.shields.io/badge/gitter-join%20chat-brightgreen.svg?style=flat-square)](https://gitter.im/labstack/echo)
[![Forum](https://img.shields.io/badge/community-forum-00afd1.svg?style=flat-square)](https://forum.labstack.com)
[![Twitter](https://img.shields.io/badge/twitter-@labstack-55acee.svg?style=flat-square)](https://twitter.com/labstack)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/labstack/echo/master/LICENSE)

## Supported Go versions

As of version 4.0.0, Echo is available as a [Go module](https://github.com/golang/go/wiki/Modules).
Therefore a Go version capable of understanding /vN suffixed imports is required:

- 1.9.7+
- 1.10.3+
- 1.11+

Any of these versions will allow you to import Echo as `github.com/labstack/echo/v4` which is the recommended
way of using Echo going forward.

For older versions, please use the latest v3 tag.

## Feature Overview
- Echo API Example

### Example

```go
package main

import (
	"github.com/labstack/echo"
	"forum-api/routers"
)

func main() {
	e := echo.New()
	routers.InitRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}
```
Routing
```go
package routers

import (
	"github.com/labstack/echo"
	handler "forum-api/handler"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/group", handler.NewGroupHandler().List())
	e.GET("/group/:identify", handler.NewGroupHandler().Show())
}
```

Handler
```go

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
```
Model
```go
package models

import "database/sql"

type Groups struct {
	ID              uint16 `json:"id"`
	Name            string `json:"name"`
	Provider        string `json:"provider"`
	ProviderID      string `json:"provider_id"`
	MemberNumber    string `json:"member_number"`
	Cover           string `json:"cover"`
	Identity        string `json:"identity"`
	DetaDescription string `json:"meta_description"`
}

func NewGroups() *Groups {
	return &Groups{}
}

func GetSQLList() string {
	return "SELECT id, name, provider, provider_id, member_number, cover, identity, meta_description FROM groups WHERE status = ?"
}

func (this *Groups) List(listRows int, status ...int) (groups []Groups, rows int64, err error) {
	var selectStmt *sql.Stmt
	DB := getConnection()
	statusType := 1;
	rows = 0;	

	if len(status) > 0 {
		statusType = status[0]
	}

	selectStmt, err = DB.Prepare(GetSQLList())	
	results, err := selectStmt.Query(statusType)
	
	for results.Next() {
		group := new(Groups)
		err = results.Scan(&group.ID, &group.Name, &group.Provider, &group.ProviderID, &group.MemberNumber, &group.Cover, &group.Identity, &group.DetaDescription)
		groups = append(groups, *group)
		rows++;
	}
	
	return groups, rows,  err
}
```