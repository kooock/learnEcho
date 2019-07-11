package admin

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func setAuth(g *echo.Group,ad Admin)  {

	g.Use(middleware.BasicAuth(func(username string, password string, context echo.Context) (b bool, e error) {
		return ad.CheckAdminInfo(username,password), nil
	}))
}
