package api

import (
	"github.com/labstack/echo"
	"learnEcho/part10/api/handlers"
)

func CookieGroup(g *echo.Group)  {
	g.GET("/main", handlers.MainCookie)
}