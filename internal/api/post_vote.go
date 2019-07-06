package api

import (
	"github.com/labstack/echo"
	apim "github.com/yukihiratype2/cookbook-server/internal/model/api"
	m "github.com/yukihiratype2/cookbook-server/internal/model/app"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// m "github.com/yukihiratype2/cookbook-server/internal/model/app"
)

type rateResp struct {
	Rate m.PostRate `json:"rate"`
}

func (ph *postHandler) Vote(c echo.Context) (err error) {
	postID, err := primitive.ObjectIDFromHex(c.Param("postID"))
	if err != nil {
		return
	}
	r := new(apim.PostVote)
	if err = c.Bind(r); err != nil {
		return
	}

	ph.postService.RatePost(postID, r.Rate)

	return
}
