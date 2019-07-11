package handlers

import (
	"github.com/labstack/echo"
	"net/http"
)

func MainAdmin(c echo.Context) error {
	return c.String(http.StatusOK,"horay you are on the secret admin main page!")
}
