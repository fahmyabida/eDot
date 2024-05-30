package middleware

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/fahmyabida/eDot/cmd/config"
	"github.com/fahmyabida/eDot/internal/app/domain"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// ErrorMiddleware is an echo middleware that handles the outgoing errors.
func AuthMiddleware(jwtConfig *config.JWTConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			nextToken, userId, err := validateToken(c, jwtConfig)
			if err != nil {
				return err
			}
			c.Response().Header().Add("token", nextToken)
			c.Response().Header().Add("user_id", userId)
			return next(c)
		}
	}
}

func validateToken(c echo.Context, jwtConfig *config.JWTConfig) (nextToken, userId string, err error) {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		err = echo.ErrUnauthorized
		return
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		err = echo.ErrUnauthorized
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtConfig.SecretKey), nil
	})

	if err != nil || !token.Valid {
		err = echo.ErrUnauthorized
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err = fmt.Errorf("unable claims jwt payload")
		return
	}

	return createToken(claims["user"], jwtConfig)
}

func createToken(user interface{}, jwtConfig *config.JWTConfig) (newToken, userId string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user,
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
	})

	newToken, err = token.SignedString([]byte(jwtConfig.SecretKey))
	if err != nil {
		err = fmt.Errorf("error signing token: %v", err)
		return
	}

	asByteJSON, err := json.Marshal(user)
	if err != nil {
		err = fmt.Errorf("unable to marshall user payload in JWT: %v", err)
		return
	}

	var userData domain.User
	if err = json.Unmarshal(asByteJSON, &userData); err != nil {
		err = fmt.Errorf("unable to unmarshall user payload in JWT: %v", err)
		return
	}

	return newToken, userData.ID, err
}
