package api

import (
	"github.com/labstack/echo"
	m "github.com/yukihiratype2/cookbook-server/internal/model/app"
	"github.com/yukihiratype2/cookbook-server/internal/service"
	// userservice "github.com/yukihiratype2/cookbook-server/internal/service/user"
)

func mountRoutes(e *echo.Echo, srv *service.Service, c m.ConfigServer) {
	v1 := e.Group("/api/v1")
	userGroup := v1.Group("/user")
	userHandler := newUserHandler(srv.UserService, c)
	authGroup := v1.Group("/auth")
	mountUserGroup(userGroup, userHandler)
	mountAuthGroup(authGroup, userHandler)
}

func mountUserGroup(userGroup *echo.Group, uh *userHandler) {
	// uh := newUserHandler(us)
	userGroup.POST("", uh.Create)
	userGroup.GET("", uh.Get)
}

func mountAuthGroup(authGroup *echo.Group, uh *userHandler) {
	authGroup.POST("/login", uh.Authentication)
}
