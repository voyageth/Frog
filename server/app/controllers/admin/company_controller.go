package admin_controllers

import (
	"github.com/revel/revel"
)

type CompanyController struct {
	AdminController
}

func (c CompanyController) Index() revel.Result {
	return c.Render()
}