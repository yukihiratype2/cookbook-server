package api

import (
	"fmt"
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
	})
	if err != nil {
		return
	}
	token, err := uh.generateToken(authParams)

	return c.JSON(http.StatusOK, echo.Map{"token": token})
}

func (uh *userHandler) generateToken(authParams *apim.AuthParams) (t string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)
	findedUser, err := uh.userService.GetUserByEmail(authParams.Email)
	if err != nil {
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	fillClams(claims, findedUser)
	t, err = token.SignedString(uh.jwtSecret)
	return
}

func fillClams(claims jwt.MapClaims, user m.User) {
	fmt.Println(user)

	claims["user"] = apim.JWTUserParams{
		User: apim.User{
			Email: user.Email,
		},
		UserName: user.UserInfo.UserName,
	}
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["iat"] = time.Now().Unix()
	claims["sub"] = user.Email
}
