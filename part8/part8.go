package main

import (
	"github.com/labstack/echo"
	"learnEcho/part8/admin"
	"learnEcho/part8/auth"
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

	cookieRouter := auth.NewAuthRouter(e,"/cookie","/jwt")

	cookieRouter.Route()

	e.Start(":8080")

}


