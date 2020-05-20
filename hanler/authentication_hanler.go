package hanler

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"strings"
)

type AuthenticationHandler struct {
}

func NewAuthenticationHandler() *AuthenticationHandler {
	return &AuthenticationHandler{}
}

func (a *AuthenticationHandler) Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		data := ctx.Request().Header["Authorization"]
		if len(data) == 0 {
			ctx.JSON(http.StatusUnauthorized, "'Authorization' is required in http request header.")
			return errors.New("'Authorization' is required in http request header")
		}
		authorization := data[0]
		if strings.HasPrefix(authorization, "Bearer ") != true {
			ctx.JSON(http.StatusUnauthorized, "Invalid 'Authorization' format. The format must be 'Authorization: Bearer [token]'")
			return errors.New("invalid 'Authorization' format. The format must be 'Authorization: Bearer [token]'")
		}
		tokenStr := authorization[7:]
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("secret_key"), nil
		})
		if claim, ok := token.Claims.(jwt.Claims); ok && token.Valid {
			ctx.Set("Authorization", claim)
			log.Println(claim)
		} else {
			ctx.JSON(500, err.Error())
			return err
		}
		return next(ctx)
	}
}
