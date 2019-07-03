package main

import (
	"github.com/labstack/echo"
	"learnEcho/part5/admin"
)




func main()  {
	e := echo.New()
	r := NewRouter(e)
	r.Route()

	adr := admin.NewAdminRouter(e,"/admin")
	ad := admin.NewAdmin("jack","1234")
	adr.SetAdminObj(ad)
	adr.Route()

	e.Start(":8080")

}


