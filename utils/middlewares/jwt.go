package middlewares

import (
	"GunTour/config"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func GenerateToken(id int, role string) string {
	claim := make(jwt.MapClaims)
	claim["authorized"] = true
	claim["role"] = role
	claim["id"] = id
	claim["expired"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	str, err := token.SignedString([]byte(config.NewConfig().JWTSecret))
	if err != nil {
		log.Error("error on token signed string", err.Error())
		return "cannot generate token"
	}

	return str
}

func ExtractToken(c echo.Context) (id int, role string) {
	token := c.Get("user").(*jwt.Token)
	if token.Valid {
		claim := token.Claims.(jwt.MapClaims)
		return int(claim["id"].(float64)), (claim["role"].(string))
	}

	return 0, ""
}
