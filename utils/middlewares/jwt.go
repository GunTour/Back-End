package middlewares

import (
	"GunTour/config"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// FUNC TO GENERATE TOKEN AFTER LOGIN
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

// FUNC TO EXTRACT TOKEN
func ExtractToken(c echo.Context) (id int, role string) {
	token := c.Get("user").(*jwt.Token)
	if token.Valid {
		claim := token.Claims.(jwt.MapClaims)
		return int(claim["id"].(float64)), string(claim["role"].(string))
	}

	return 0, ""
}
