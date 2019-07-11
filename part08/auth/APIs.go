package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

func mainCookie(c echo.Context) error{
	return c.String(http.StatusOK,"you are on the secret cookie page!")
}

func mainJwt(c echo.Context) error{
	user := c.Get("user")
	token := user.(*jwt.Token)

	claims := token.Claims.(jwt.MapClaims)

	log.Println("User Name: ", claims["name"], "User ID: ",claims["jti"])

	return c.String(http.StatusOK,"you are on the secret jwt page!")
}
