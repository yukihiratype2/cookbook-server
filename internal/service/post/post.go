package postservice

import (
	"github.com/yukihiratype2/cookbook-server/internal/db"
)

type PostService struct {
	PostDB *db.PostHandler
}

func New() *PostService {
	ps := PostService{}
	return &ps
}
