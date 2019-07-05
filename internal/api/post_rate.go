package api

import (
	"github.com/labstack/echo"
	apim "github.com/yukihiratype2/cookbook-server/internal/model/api"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// m "github.com/yukihiratype2/cookbook-server/internal/model/app"
)

func (ph *postHandler) Rate(c echo.Context) (err error) {
	postID, err := primitive.ObjectIDFromHex(c.Param("postID"))
	if err != nil {
		return
	}
	r := new(apim.PostRate)
	if err = c.Bind(r); err != nil {
		return
	}

	ph.postService.RatePost(postID, r.Rate)

	return
}
