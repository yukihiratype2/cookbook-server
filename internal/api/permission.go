package api

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type customCliams struct {
	jwt.StandardClaims
}

func authWithJWT(secret string) echo.MiddlewareFunc {
	// TODO
	config := middleware.JWTConfig{
		SigningKey: []byte(secret),
	}
	return middleware.JWTWithConfig(config)
}
