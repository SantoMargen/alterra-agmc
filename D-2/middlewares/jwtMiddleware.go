package middlewares

import (
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func CreateToken(userId uint) (string, error) {
	Claims := jwt.MapClaims{}
	Claims["authorized"] = true
	Claims["userId"] = userId
	Claims["exp"] = time.Now().Add(time.Hour * 10).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

type TokenData struct {
	UserId       uint
	IsAuthorized bool
}

func ExtractTokenData(e echo.Context) int {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		usrId := claims["userId"].(int)
		return usrId
	}
	return 0
}
