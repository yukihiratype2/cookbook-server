package api

import (
	"net/http"

	"github.com/labstack/echo"
	apim "github.com/yukihiratype2/cookbook-server/internal/model/api"
)

type postHandler struct {
}

func newPostHandler() *postHandler {
	ph := postHandler{}
	return &ph
}

func (ph *postHandler) Create(c echo.Context) (err error) {
	u := new(apim.CreatePostParams)
	if err = c.Bind(u); err != nil {
		return
	}
	c.String(http.StatusOK, "post")
	return
}
