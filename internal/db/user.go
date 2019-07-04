package db

import (
	"context"
	m "github.com/yukihiratype2/cookbook-server/internal/model/app"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

// UserHandler struct
type UserHandler struct {
	UserCollection *mongo.Collection
}

type UserHandlerMethod interface {
	Create()
}

func newUserHandler(uc *mongo.Collection) *UserHandler {
	uh := UserHandler{
		UserCollection: uc,
	}
	return &uh
}

func (uh UserHandler) Create(newUser *m.User) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = uh.UserCollection.InsertOne(ctx, newUser)
	return
}

func (uh UserHandler) Update(userToUpdate *m.User) (err error) {
	return
}

func (uh UserHandler) Delete(userToDelete *m.User) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = uh.UserCollection.DeleteOne(ctx, userToDelete)
	return
}

func (uh UserHandler) FindByEmail(email string) (findedUser m.User, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = uh.UserCollection.FindOne(ctx, &bson.M{"email": email}).Decode(&findedUser)
	return
}
