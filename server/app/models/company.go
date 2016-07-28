package models

import (
	"github.com/revel/revel"
)

type Company struct {
	No   int
	Name string
}

func (company *Company) Validate(v *revel.Validation) {
	v.Check(
		company.Name,
		revel.Required{},
		revel.MinSize{1},
		revel.MaxSize{200},
	)
}
