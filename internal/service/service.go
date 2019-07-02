package service

import (
	"github.com/yukihiratype2/cookbook-server/internal/db"
	userservice "github.com/yukihiratype2/cookbook-server/internal/service/user"
)

type Service struct {
	UserService *userservice.UserService
}

func New(dh *db.DB) *Service {
	s := Service{}
	s.UserService = userservice.New(dh.UserHandler)
	return &s
}
