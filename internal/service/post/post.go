package postservice

import (
	"github.com/yukihiratype2/cookbook-server/internal/db"
	m "github.com/yukihiratype2/cookbook-server/internal/model/app"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
