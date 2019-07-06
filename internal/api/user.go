package api

import (
	"fmt"
	"net/http"
	"time"

	// "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	apim "github.com/yukihiratype2/cookbook-server/internal/model/api"
	m "github.com/yukihiratype2/cookbook-server/internal/model/app"
	userservice "github.com/yukihiratype2/cookbook-server/internal/service/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/go-playground/validator.v9"
)

type userHandler struct {
	userService *userservice.UserService
	validator   *validator.Validate
	jwtSecret   []byte
}

func newUserHandler(us *userservice.UserService, c m.ConfigServer) *userHandler {
	uh := userHandler{}
	uh.userService = us
	uh.validator = validator.New()
	uh.jwtSecret = []byte(c.Secret)
	return &uh
}

func (uh *userHandler) Create(c echo.Context) (err error) {
	u := new(apim.CreateUserParams)
	if err = c.Bind(u); err != nil {
		return
	}
	if err = uh.validator.Struct(u); err != nil {
		return
	}
	uh.userService.Create(&m.User{
		ID:       primitive.NewObjectID(),
		Email:    u.Email,
		Password: u.Password,
		UserInfo: m.UserInfo{
			UserName:    u.UserName,
			CreatedAt:   time.Now(),
			LastLogin:   time.Now(),
			LastLoginAt: c.RealIP(),
		},
		Role: []string{"USER"},
	})
	return c.JSON(http.StatusOK, echo.Map{"message": "Successful create user"})
}

func (uh *userHandler) Get(c echo.Context) (err error) {
	cc := c.(*Context)
	fmt.Printf("%+v\n", cc.jwtPaylaod)

	// claims := user.Claims.(*jwtCustomClaims)

	c.String(http.StatusOK, "wr")
	return
}

func (uh *userHandler) Delete(c echo.Context) {

}
