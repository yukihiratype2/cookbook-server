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

func (ph *postHandler) Get(c echo.Context) (err error) {
	postID := c.Param("postID")
	findedPost, err := ph.postService.GetPostByID(postID)

	if err != nil {
		return
	}
	c.JSON(http.StatusOK, &findedPost)
	return
}

func mountPostGroup(postGroup *echo.Group, ph *postHandler) {
	postGroup.POST("", ph.Create)
	postGroup.GET("/:postID", ph.Get)
	postGroup.POST("/:postID/rate", ph.Rate)
}
