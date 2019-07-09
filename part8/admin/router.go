package admin

import "github.com/labstack/echo"

type Router interface {
	SetAdminObj(Admin)
	SetAdminPoint(*echo.Echo, string)
	Route()
}

type router struct {
	g *echo.Group
	ad Admin
}


func NewAdminRouter(e *echo.Echo, adminPoint string) Router{
	r := &router{}
	r.SetAdminPoint(e,adminPoint)
	return r
}

func (r *router)SetAdminObj(ad Admin){
	ad.SetAdminGroup(r.g)
}

func (r *router)SetAdminPoint(e *echo.Echo,point string){
	r.g = e.Group(point)
}


func (r *router)Route(){
	r.g.GET("/main", mainAdmin)
}