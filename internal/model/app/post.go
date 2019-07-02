package model

import (
	"time"
)

type Post struct {
	Title          string      `bson:"title"`
	PostInfo       PostInfo    `bson:"post_info"`
	Rate           [5]int      `bson:"rate"`
	DefaultFormula int         `bson:"default_formula"`
	Formulas       []Formula   `bson:"formulas"`
	Cuisine        PostCuisine `bson:"cuisine"`
	Private        bool        `bson:"private"`
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
	Describe string `bson:"describe"`
	Example  string `bson:"example"`
}

type FormulaMedia struct {
	MediaList []string
}
