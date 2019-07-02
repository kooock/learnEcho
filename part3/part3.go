package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"io/ioutil"
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

	return c.JSON(http.StatusBadRequest, NewErrorMsg("you need to lets us know if you want json or string data"))

}

func addCat(c echo.Context) error {

	cat := CreateCat("","")
	b,err := ioutil.ReadAll(c.Request().Body)
	defer c.Request().Body.Close()
	if err != nil {
		msg := fmt.Sprintf("Failed reading the request body for addCat: %s",err)
		c.Logger().Error(msg)
		return c.JSON(http.StatusInternalServerError, NewErrorMsg(msg))
	}

	err = json.Unmarshal(b,cat)
	if err != nil {
		msg := fmt.Sprintf("Failed unmarshaling in addCat: %s",err)
		c.Logger().Error(msg)
		return c.JSON(http.StatusInternalServerError, NewErrorMsg(msg))
	}

	c.Logger().Info(fmt.Sprintf("to add cat is success: %v",cat))
	return c.String(http.StatusOK,fmt.Sprintf("the cat is added : %v",cat))
}

func addDog(c echo.Context) error {
	dog := CreateDog("","")
	err := c.Bind(dog)
	if err != nil {
		c.Logger().Error("failed processing addDog request: %s",err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	c.Logger().Info(fmt.Sprintf("to add cat is success: %v",dog))
	return c.String(http.StatusOK,fmt.Sprintf("the cat is added : %v",dog))
}

func addHamster(c echo.Context) error {
	hamster := CreateDog("","")
	err := c.Bind(hamster)
	if err != nil {
		c.Logger().Error("failed processing addDog request: %s",err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	c.Logger().Info(fmt.Sprintf("to add cat is success: %v",hamster))
	return c.String(http.StatusOK,fmt.Sprintf("the cat is added : %v",hamster))
}

func main()  {

	e := echo.New()
	r := NewRouter(e)
	r.Route()
	e.Start(":8080")

}


