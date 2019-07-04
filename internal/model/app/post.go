package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Post struct {
	ID             primitive.ObjectID `bson:"_id"`
	Title          string             `bson:"title"`
	PostInfo       PostInfo           `bson:"post_info"`
	Rate           [5]int             `bson:"rate"`
	DefaultFormula int                `bson:"default_formula"`
	Formula        []Formula          `bson:"formula"`
	Cuisine        PostCuisine        `bson:"cuisine"`
	Private        bool               `bson:"private"`
}

type PostCuisine struct{}

type PostInfo struct {
	CreatedAt  time.Time `bson:"create_at"`
	LastModify time.Time `bons:"last_modify"`
}

type Formula struct {
	Media           FormulaMedia   `bson:"media"`
	Describe        string         `bson:"describe"`
	FormulaContent  FormulaContent `bson:"formula_content"`
	FormulaTitle    string         `bson:"formula_title"`
	PostFormulaInfo FormulaInfo    `bson:"formula_info"`
}

type FormulaInfo struct {
	CreatedAt  time.Time `bson:"create_at"`
	LastModify time.Time `bons:"last_modify"`
}

type FormulaContent struct {
	Material []FormulaMaterial `bson:"material"`
	Step     []FormulaStep     `bson:"step"`
	Note     string            `bson:"note"`
}

type FormulaMaterial struct {
	Name  string `bson:"name"`
	Usage string `bson:"usage"`
	Note  string `bson:"note"`
}

type FormulaStep struct {
	Describe string    `bson:"describe"`
	Example  PostAsset `bson:"example"`
}

type FormulaMedia struct {
	Picture []PostAsset `bson:"picture"`
}

type PostAsset struct {
	AssetURL string `bson:"asset_url"`
}
