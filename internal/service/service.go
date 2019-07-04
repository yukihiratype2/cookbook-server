package service

import (
	"github.com/yukihiratype2/cookbook-server/internal/db"
	postservice "github.com/yukihiratype2/cookbook-server/internal/service/post"
	userservice "github.com/yukihiratype2/cookbook-server/internal/service/user"
)

type Service struct {
	UserService *userservice.UserService
	PostService *postservice.PostService
}

func New(dh *db.DB) *Service {
	s := Service{}
	s.UserService = userservice.New(dh.UserHandler)
	s.PostService = postservice.New(dh.PostHandler)
	return &s
}
