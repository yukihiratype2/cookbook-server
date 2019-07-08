package postservice

import (
	m "github.com/yukihiratype2/cookbook-server/internal/model/app"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Vote for post
func (ps *PostService) Vote(postID primitive.ObjectID, vote int) (rate m.PostRate, err error) {
	return ps.PostDB.Vote(postID, vote)
}
