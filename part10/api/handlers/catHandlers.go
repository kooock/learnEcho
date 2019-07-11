package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"
	"learnEcho/part10/pet"

	errs "learnEcho/part10/errors"
)

func GetCats(c echo.Context) error {
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
		return c.JSON(http.StatusOK, pet.CreateCat(catName,catType))
	}

	return c.JSON(http.StatusBadRequest, errs.NewErrorMsg("you need to lets us know if you want json or string data"))

}

func AddCat(c echo.Context) error {

	cat := pet.CreateCat("","")
	b,err := ioutil.ReadAll(c.Request().Body)
	defer c.Request().Body.Close()
	if err != nil {
		msg := fmt.Sprintf("Failed reading the request body for addCat: %s",err)
		c.Logger().Error(msg)
		return c.JSON(http.StatusInternalServerError, errs.NewErrorMsg(msg))
	}

	err = json.Unmarshal(b,cat)
	if err != nil {
		msg := fmt.Sprintf("Failed unmarshaling in addCat: %s",err)
		c.Logger().Error(msg)
		return c.JSON(http.StatusInternalServerError, errs.NewErrorMsg(msg))
	}

	c.Logger().Info(fmt.Sprintf("to add cat is success: %v",cat))
	return c.String(http.StatusOK,fmt.Sprintf("the cat is added : %v",cat))
}

