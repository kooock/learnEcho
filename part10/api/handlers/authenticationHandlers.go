package handlers

import (
	"github.com/labstack/echo"
	"learnEcho/part10/auth"
	"learnEcho/part10/user"
	"log"
	"net/http"
	"time"
)

func Login(c echo.Context) error {
	username := c.QueryParam("username")
	password := c.QueryParam("password")

	u := user.Get()

	if u.CheckLogin(username,password) {
		cookie := new(http.Cookie)

		cookie.Name = "sessionID"
		cookie.Value = "some_string"
		cookie.Expires = time.Now().Add(48 * time.Hour)

		c.SetCookie(cookie)

		//create jwt token
		token,err := auth.CreateJwtToken()

		if err != nil {
			log.Println("Error create JWT token",err)
			return c.String(http.StatusInternalServerError,"something went wrong")
		}

		jwtCookie := new(http.Cookie)

		jwtCookie.Name = "JWTCookie"
		jwtCookie.Value = token
		jwtCookie.Expires = time.Now().Add(48 * time.Hour)

		c.SetCookie(jwtCookie)



		return c.JSON(http.StatusOK, map[string]string{
			"message" : "You were logged in",
			"token" : token,
		})
	}

	return c.String(http.StatusUnauthorized,"YOur username or password were wrong")


}




