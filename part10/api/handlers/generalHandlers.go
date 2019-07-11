package handlers

import (
	"github.com/labstack/echo"
	"net/http"
)

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello world")
}
