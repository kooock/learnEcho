package cookie

import (
	"github.com/labstack/echo"
	"net/http"
	"strings"
)

func checkCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("sessionID")

		if strings.Contains(err.Error(),"named auth not present") {
			return c.String(http.StatusUnauthorized,"you dont have any auth")
		}


		if err != nil {
			return err
		}

		if cookie.Value == "some_string" {
			return next(c)
		}

		return c.String(http.StatusUnauthorized,"you dont have the right auth, auth")
	}

}