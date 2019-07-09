package auth

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JwtClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

func CreateJwtToken() (string,error){
	claims := JwtClaims{
		"jack",
		jwt.StandardClaims{
			Id: "main_user_id",
			ExpiresAt: time.Now().Add(24*time.Hour).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512,claims)

	token, err := rawToken.SignedString([]byte("mySecret"))
	if err != nil {
		return "",err
	}

	return token, nil
}