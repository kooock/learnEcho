package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func SetJwtMiddlewares(g *echo.Group)  {
	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "H5512",
		SigningKey: []byte("mySecret"),
		TokenLookup: "cookie:JWTCookie",
	}))
}
