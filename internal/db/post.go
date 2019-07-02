package db

import (
	"context"
	m "github.com/yukihiratype2/cookbook-server/internal/model/app"
	// "go.mongodb.org/mongo-driver/bson"
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
