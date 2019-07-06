package api

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	apim "github.com/yukihiratype2/cookbook-server/internal/model/api"
	m "github.com/yukihiratype2/cookbook-server/internal/model/app"
	"net/http"
	"time"
)

func (uh *userHandler) Authentication(c echo.Context) (err error) {
	authParams := new(apim.AuthParams)
	if err = c.Bind(authParams); err != nil {
		return
	}
	if err = uh.validator.Struct(authParams); err != nil {
		return
	}
	err = uh.userService.Authentication(&m.User{
		Email:    authParams.Email,
		Password: authParams.Password,
		UserInfo: m.UserInfo{
			LastLogin:   time.Now(),
			LastLoginAt: c.RealIP(),
		},
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, echo.Map{"message": "email or password invalid"})
		return
	}
	token, err := uh.generateToken(authParams)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, echo.Map{"token": token})
	return
}

func (uh *userHandler) generateToken(authParams *apim.AuthParams) (t string, err error) {
	findedUser, err := uh.userService.GetUserByEmail(authParams.Email)
	claims := fillClams(findedUser)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if err != nil {
		return
	}
	t, err = token.SignedString(uh.jwtSecret)
	return
}

func fillClams(user m.User) *jwtCustomClaims {
	claims := jwtCustomClaims{}
	claims.User = apim.JWTUserParams{
		User: apim.User{
			Email: user.Email,
		},
		UserName: user.UserInfo.UserName,
	}
	claims.ExpiresAt = time.Now().Add(time.Hour * 24).Unix()
	claims.IssuedAt = time.Now().Unix()
	claims.Subject = user.Email
	return &claims
}
