package postservice

import (
	"github.com/yukihiratype2/cookbook-server/internal/db"
	m "github.com/yukihiratype2/cookbook-server/internal/model/app"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
)

type PostService struct {
	PostDB *db.PostHandler
}

func New(ph *db.PostHandler) *PostService {
	ps := PostService{}
	ps.PostDB = ph
	return &ps
}

func (ps *PostService) Create(newPost *m.Post) {
	ps.PostDB.Create(newPost)
}

func (ps *PostService) GetPostByID(postID string) (findedPost *m.Post, err error) {
	postIDInObject, err := primitive.ObjectIDFromHex(postID)
	postToFind := m.Post{
		ID: postIDInObject,
	}
	return ps.PostDB.Get(&postToFind)
}

//RatePost accept a int
func (ps *PostService) RatePost(postID primitive.ObjectID, rate int) (rateAfterVote m.PostRate, err error) {
	err = ps.PostDB.Update(postID, &bson.D{{"$inc", bson.D{
		{"rate." + strconv.Itoa(rate), 1},
	}}})
	if err != nil {
		return
	}

	return
}
