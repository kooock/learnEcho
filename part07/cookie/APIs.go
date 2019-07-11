package cookie

import (
	"github.com/labstack/echo"
	"net/http"
)

func mainCookie(c echo.Context) error{
	return c.String(http.StatusOK,"you are on the secret auth page!")
}

