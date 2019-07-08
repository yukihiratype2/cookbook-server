package db

import (
	m "github.com/yukihiratype2/cookbook-server/internal/model/app"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ph *PostHandler) CreateFormula(postID primitive.ObjectID, newFormula m.Formula) (err error) {
	return
}
