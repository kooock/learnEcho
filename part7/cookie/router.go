package cookie

import (
	"github.com/labstack/echo"
)

type Router interface {
	SetCookiePoint(*echo.Echo, string)
	Route()
}

type router struct {
	g *echo.Group
}


func NewCookieRouter(e *echo.Echo, cookiePoint string) Router{
	r := &router{}
	r.SetCookiePoint(e,cookiePoint)
	r.g.Use(checkCookie)
	return r
}

func (r *router)SetCookiePoint(e *echo.Echo,point string){
	r.g = e.Group(point)
}


func (r *router)Route(){
	r.g.GET("/main", mainCookie)
}