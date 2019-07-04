package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	apim "github.com/yukihiratype2/cookbook-server/internal/model/api"
	m "github.com/yukihiratype2/cookbook-server/internal/model/app"
	postservice "github.com/yukihiratype2/cookbook-server/internal/service/post"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type postHandler struct {
	postService *postservice.PostService
}

func newPostHandler(ps *postservice.PostService) *postHandler {
	ph := postHandler{}
	ph.postService = ps
	return &ph
}

func (ph *postHandler) Create(c echo.Context) (err error) {
	p := new(apim.CreatePostParams)
	if err = c.Bind(p); err != nil {
		return
	}

	fmt.Printf("%+v\n", p)

	ph.postService.Create(&m.Post{
		// TODO fix pointer
		ID:    primitive.NewObjectID(),
		Title: p.Title,
		Formula: []m.Formula{m.Formula{
			FormulaTitle: p.Formula[0].FormulaTitle,
			Describe:     p.Formula[0].Describe,
		}},
	})
	c.String(http.StatusOK, "post")
	return
}

func mountPostGroup(postGroup *echo.Group, ph *postHandler) {
	postGroup.POST("", ph.Create)
}
