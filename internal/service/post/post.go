package postservice

import (
	"github.com/yukihiratype2/cookbook-server/internal/db"
	m "github.com/yukihiratype2/cookbook-server/internal/model/app"
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
