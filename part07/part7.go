package main

import (
	"github.com/labstack/echo"
	"learnEcho/part07/admin"
	"learnEcho/part07/cookie"
)




func main()  {
	e := echo.New()

	UseServerHeaer(e)


	r := NewRouter(e)
	r.Route()

	adminRouter := admin.NewAdminRouter(e,"/admin")
	ad := admin.NewAdmin("jack","1234")
	adminRouter.SetAdminObj(ad)
	adminRouter.Route()

	cookieRouter := cookie.NewCookieRouter(e,"/auth")

	cookieRouter.Route()

	e.Start(":8080")

}


