package db

import (
	"context"
	"time"

	m "github.com/yukihiratype2/cookbook-server/internal/model/app"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB struct
type DB struct {
	DBClient    *mongo.Client
	UserHandler *UserHandler
	PostHandler *PostHandler
}

// NewDbInstance struct
func New(c m.ConfigDB) *DB {
	db := DB{}
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(c.DBAddress))
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}
	dbInUse := client.Database(c.DBName)

	db.DBClient = client
	db.UserHandler = newUserHandler(dbInUse.Collection("user"))
	db.PostHandler = newPostHandler(dbInUse.Collection("post"))
	return &db
}
