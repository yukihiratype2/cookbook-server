package api

import (
	"github.com/labstack/echo"
	m "github.com/yukihiratype2/cookbook-server/internal/model/app"
	"github.com/yukihiratype2/cookbook-server/internal/service"
)

// Start Server
func Start(c m.ConfigServer, srv *service.Service) {
	e := echo.New()
	mountMiddlewale(e, c)

	mountRoutes(e, srv, c)
	e.Logger.Fatal(e.Start(":" + c.ServerPort))
}
