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

func (r *router)RouteAdminGroup(groupPoint string) error{
	if 	r.e == nil{
		return errors.New("not found routing framework")
	}
	g := r.e.Group(groupPoint)
	g.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))
	g.GET("/main",mainAdmin)

	return nil
}

func (r *router)Route() error{
	if 	r.e == nil{
		return errors.New("not found routing framework")
	}

	r.RouteAdminGroup("/admin")

	r.e.GET("/", hello)
	r.e.GET("/cats/:data",getCats)
	r.e.POST("/cats",addCat)
	r.e.POST("/dogs",addDog)
	r.e.POST("/hamsters",addHamster)

	return nil
}