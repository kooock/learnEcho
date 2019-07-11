package main

import (
	"errors"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Router interface {
	Route() error
}

type router struct {
	e *echo.Echo

}

func NewRouter(e *echo.Echo) Router{
	return &router{e:e}
}

func (r *router)Route() error{
	if 	r.e == nil{
		return errors.New("not found routing framework")
	}

	r.e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:"./part09/static",
	}))

	r.e.GET("/hello", hello)
	r.e.GET("/cats/:data",getCats)
	r.e.GET("/login",login)
	r.e.POST("/cats",addCat)
	r.e.POST("/dogs",addDog)
	r.e.POST("/hamsters",addHamster)


	return nil
}
