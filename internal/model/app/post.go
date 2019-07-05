package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Post struct {
	ID             primitive.ObjectID `json:"_id" bson:"_id"`
	Title          string             `json:"titlel" bson:"title"`
	PostInfo       PostInfo           `json:"post_info" bson:"post_info"`
	Rate           PostRate           `json:"rate" bson:"rate"`
	DefaultFormula int                `json:"default_formula" bson:"default_formula"`
	Formula        []Formula          `json:"formula" bson:"formula"`
	Cuisine        PostCuisine        `json:"cuisine" bson:"cuisine"`
	Private        bool               `json:"private" bson:"private"`
	Viewed         int                `json:"viewed" bson:"viewed"`
}

type PostRate [5]int

type PostCuisine struct{}

type PostInfo struct {
	CreatedAt  time.Time `json:"create_at" bson:"create_at"`
	LastModify time.Time `json:"last_modify" bons:"last_modify"`
}

type Formula struct {
	Media           FormulaMedia   `json:"medial" bson:"media"`
	Describe        string         `json:"describe" bson:"describe"`
	FormulaContent  FormulaContent `json:"formula_content" bson:"formula_content"`
	FormulaTitle    string         `json:"formula_title" bson:"formula_title"`
	PostFormulaInfo FormulaInfo    `json:"formula_info" bson:"formula_info"`
}

type FormulaInfo struct {
	CreatedAt  time.Time `json:"create_at" bson:"create_at"`
	LastModify time.Time `json:"last_modify" bons:"last_modify"`
}

type FormulaContent struct {
	Material []FormulaMaterial `json:"material" bson:"material"`
	Step     []FormulaStep     `json:"step" bson:"step"`
	Note     string            `json:"note" bson:"note"`
}

type FormulaMaterial struct {
	Name  string `json:"name" bson:"name"`
	Usage string `json:"usage" bson:"usage"`
	Note  string `json:"note" bson:"note"`
}

type FormulaStep struct {
	Describe string    `json:"describe" bson:"describe"`
	Example  PostAsset `json:"example" bson:"example"`
}

type FormulaMedia struct {
	Picture []PostAsset `json:"picture" bson:"picture"`
}

type PostAsset struct {
	AssetURL string `json:"asset_url" bson:"asset_url"`
}
