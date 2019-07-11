package handlers

import (
	"fmt"
	"github.com/labstack/echo"
	"learnEcho/part10/pet"
	"net/http"
)

func AddDog(c echo.Context) error {
	dog := pet.CreateDog("","")
	err := c.Bind(dog)
	if err != nil {
		c.Logger().Error("failed processing addDog request: %s",err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	c.Logger().Info(fmt.Sprintf("to add cat is success: %v",dog))
	return c.String(http.StatusOK,fmt.Sprintf("the cat is added : %v",dog))
}

