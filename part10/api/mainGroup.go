package api

import (
	"github.com/labstack/echo"
	"learnEcho/part10/api/handlers"
)

func MainGroup(e *echo.Echo)  {
	e.GET("/hello", handlers.Hello)
	e.GET("/cats/:data", handlers.GetCats)
	e.GET("/login", handlers.Login)
	e.POST("/cats", handlers.AddCat)
	e.POST("/dogs", handlers.AddDog)
	e.POST("/hamsters", handlers.AddHamster)

}
