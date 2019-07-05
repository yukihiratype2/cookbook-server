package model

import (
	m "github.com/yukihiratype2/cookbook-server/internal/model/app"
)

type Post struct {
	Title          string      `json:"title"`
	Rate           [5]int      `json:"rate"`
	DefaultFormula int         `json:"default_formula"`
	Formula        []m.Formula `json:"formula"`
	Cuisine        PostCuisine `json:"cuisine"`
}

type PostRate struct {
	Rate int `json:"rate"`
}

type PostCuisine struct{}

type Formula struct {
	Describe       string         `json:"describe"`
	FormulaTitle   string         `json:"title"`
	FormulaContent FormulaContent `json:"content"`
}

type FormulaContent struct {
	Note     string              `json:"note"`
	Material []m.FormulaMaterial `json:"material"`
	Step     []m.FormulaStep     `json:"step"`
}

type FormulaStep struct {
	Describe string `json:"describe"`
}

type CreatePostParams struct {
	Post
}
