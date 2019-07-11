package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)



func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello world")
}

func getCats(c echo.Context) error {
	catName := c.QueryParam("name")
	catType := c.QueryParam("type")

	dataType := c.Param("data")

	switch dataType {
	case "string":
		return c.String(http.StatusOK,fmt.Sprintf(
			"your cat name is %s\n" +
				"and his type is %s\n",
			catName,catType))
	case "json":
		return c.JSON(http.StatusOK, CreateCat(catName,catType))
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"error" : "you need to lets us know if you want json or string data",
	})

}

func main()  {

	e := echo.New()

	e.GET("/", hello)
	e.GET("/cats/:data",getCats)
	e.Start(":8080")

}


