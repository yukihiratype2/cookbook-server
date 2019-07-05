package trans

import (
	apim "github.com/yukihiratype2/cookbook-server/internal/model/api"
	m "github.com/yukihiratype2/cookbook-server/internal/model/app"
)

func UserToAPI(u *m.User) {

}

func UserToInternal(u *apim.User) {}

func PostToAPI(p *m.Post) *apim.Post {
	apiPost := apim.Post{}
	apiPost.Title = p.Title
	apiPost.Rate = p.Rate
	apiPost.DefaultFormula = p.DefaultFormula

	for _, pFormula := range p.Formula {
		apiFormula := apim.Formula{
			Describe:     pFormula.Describe,
			FormulaTitle: pFormula.FormulaTitle,
			FormulaContent: apim.FormulaContent{
				Note: pFormula.FormulaContent.Note,
			},
		}

		apiPost.Formula = append(apiPost.Formula, apiFormula)
	}

	return &apiPost
}

func PostToInternal(u apim.Post) {

}
