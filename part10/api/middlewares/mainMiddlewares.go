package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func SetMainMiddlewares(e *echo.Echo)  {
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:"./part9/static",
	}))

	e.Use(ServerHeader)
}

func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "koock/1.0")
		c.Response().Header().Set("notReallyHeader", "thisHaveNoMeaning ")
		return next(c)
	}
}