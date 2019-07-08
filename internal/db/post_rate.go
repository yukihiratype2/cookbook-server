package db

import (
	"context"
	"time"

	m "github.com/yukihiratype2/cookbook-server/internal/model/app"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
)

type queryPostWithID struct {
	ID primitive.ObjectID `bson:"_id"`
}

// Vote accept rate and return rate
func (ph *PostHandler) Vote(postID primitive.ObjectID, vote int) (rate m.PostRate, err error) {
	var Before options.ReturnDocument = 1
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	postFilter := queryPostWithID{postID}
	postUpdated := new(m.Post)
	err = ph.PostCollection.FindOneAndUpdate(ctx, &postFilter, &bson.M{"$inc": bson.M{
		"rate." + strconv.Itoa(vote): 1,
	}}, &options.FindOneAndUpdateOptions{
		ReturnDocument: &Before,
	}).Decode(&postUpdated)
	rate = postUpdated.Rate

	return
}
