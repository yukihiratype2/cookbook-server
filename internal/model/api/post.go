package model

import ()

type Post struct {
	Title          string      `json:"title"`
	Rate           [5]int      `json:"rate"`
	DefaultFormula int         `json:"default_formula"`
	Formula        []Formula   `json:"formula"`
	Cuisine        PostCuisine `json:"cuisine"`
}

type PostCuisine struct{}

type Formula struct {
	Describe       string         `json:"describe"`
	FormulaTitle   string         `json:"title"`
	FormulaContent FormulaContent `json:"content"`
}

type FormulaContent struct {
	Note     string            `json:"note"`
	Material []FormulaMaterial `json:"material"`
	Step     FormulaStep       `json:"step"`
}

type FormulaStep struct {
	Describe string `json:"describe"`
}
type FormulaMaterial struct {
	Name  string `json:"name"`
	Usage string `json:"usage"`
	Note  string `json:"note"`
}

type CreatePostParams struct {
	Post
}
