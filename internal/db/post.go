package db

import (
	"context"
	m "github.com/yukihiratype2/cookbook-server/internal/model/app"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

// PostHandler struct
type PostHandler struct {
	PostCollection *mongo.Collection
}

func newPostHandler(pc *mongo.Collection) *PostHandler {
	ph := PostHandler{
		PostCollection: pc,
	}
	return &ph
}

func (ph *PostHandler) Create(newPost *m.Post) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = ph.PostCollection.InsertOne(ctx, newPost)
	return
}

func (ph *PostHandler) Get(postToFind *m.Post) (findedPost *m.Post, err error) {
	// findedPost := new(m.Post)
	// TODO
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = ph.PostCollection.FindOne(ctx, &bson.M{"_id": postToFind.ID}).Decode(&findedPost)
	return
}
