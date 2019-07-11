package auth

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Router interface {
	SetCookiePoint(*echo.Echo, string)
	SetJwtPoint(*echo.Echo, string)
	Route()
}

type router struct {
	cookieGroup *echo.Group
	jwtGroup *echo.Group
}


func NewAuthRouter(e *echo.Echo, cookiePoint string,jwtPoint string) Router{
	r := &router{}

	r.SetCookiePoint(e,cookiePoint)
	r.cookieGroup.Use(checkCookie)

	r.SetJwtPoint(e,jwtPoint)
	r.jwtGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "H5512",
		SigningKey: []byte("mySecret"),
		TokenLookup: "cookie:JWTCookie",
	}))
	return r
}

func (r *router)SetCookiePoint(e *echo.Echo,point string){
	r.cookieGroup = e.Group(point)
}

func (r *router)SetJwtPoint(e *echo.Echo,point string){
	r.jwtGroup = e.Group(point)
}


func (r *router)Route(){
	r.cookieGroup.GET("/main", mainCookie)


	r.jwtGroup.GET("/main", mainJwt)
}