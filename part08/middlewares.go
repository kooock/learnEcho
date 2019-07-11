package main

import "github.com/labstack/echo"

func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "koock/1.0")
		c.Response().Header().Set("notReallyHeader", "thisHaveNoMeaning ")
		return next(c)
	}
}

func UseServerHeaer(e *echo.Echo){
	e.Use(ServerHeader)
}