package api

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	apim "github.com/yukihiratype2/cookbook-server/internal/model/api"
	m "github.com/yukihiratype2/cookbook-server/internal/model/app"
)

type jwtCustomClaims struct {
	jwt.StandardClaims
	User apim.JWTUserParams `json:"user"`
}

type jwtWhitelist struct {
	path   string
	method string
}

func mountMiddlewale(e *echo.Echo, c m.ConfigServer) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(jwtMiddleware(c.Secret))
}

func jwtMiddleware(token string) echo.MiddlewareFunc {
	jwtConfig := middleware.JWTConfig{
		Skipper:    jwtSkipper,
		SigningKey: []byte(token),
		Claims:     &jwtCustomClaims{},
	}
	return middleware.JWTWithConfig(jwtConfig)
}

func jwtSkipper(c echo.Context) bool {
	path := jwtWhitelist{
		c.Path(),
		c.Request().Method,
	}

	whitelist := [...]jwtWhitelist{jwtWhitelist{"/api/v1/auth/login", "POST"}, jwtWhitelist{"/api/v1/user", "POST"}}
	for _, whitepath := range whitelist {
		if path == whitepath {
			return true
		}
	}
	return false
}
