package router

import (
	"github.com/labstack/echo"
	"learnEcho/part10/api"
	"learnEcho/part10/api/middlewares"
)

func New() *echo.Echo {
	e := echo.New()

	// create groups
	adminGroup := e.Group("/admin")
	cookieGroup := e.Group("/cookie")
	jwtGroup := e.Group("/jwt")

	//set all middlewares
	middlewares.SetMainMiddlewares(e)
	middlewares.SetAdminMiddlewares(adminGroup)
	middlewares.SetCookieMiddlewares(cookieGroup)
	middlewares.SetJwtMiddlewares(jwtGroup)

	api.MainGroup(e)
	api.AdminGroup(adminGroup)
	api.CookieGroup(cookieGroup)
	api.JwtGroup(jwtGroup)

	return e

}

