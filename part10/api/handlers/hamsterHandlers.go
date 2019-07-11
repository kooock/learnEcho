package handlers

import (
	"fmt"
	"github.com/labstack/echo"
	"learnEcho/part10/pet"
	"net/http"
)

func AddHamster(c echo.Context) error {
	hamster := pet.CreateDog("","")
	err := c.Bind(hamster)
	if err != nil {
		c.Logger().Error("failed processing addDog request: %s",err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	c.Logger().Info(fmt.Sprintf("to add cat is success: %v",hamster))
	return c.String(http.StatusOK,fmt.Sprintf("the cat is added : %v",hamster))
}

