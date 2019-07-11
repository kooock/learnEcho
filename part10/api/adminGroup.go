package api

import (
	"github.com/labstack/echo"
	"learnEcho/part10/api/handlers"
)

func AdminGroup(g *echo.Group)  {
	g.GET("/main", handlers.MainAdmin)
}
