package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"learnEcho/part10/admin"
)

func SetAdminMiddlewares(g *echo.Group)  {


	adminObj := admin.New("jack","1234")


	g.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	g.Use(middleware.BasicAuth(func(username string, password string, context echo.Context) (b bool, e error) {
		return adminObj.CheckAdminInfo(username,password), nil
	}))
}
