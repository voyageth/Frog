package admin_controllers

import (
	"github.com/revel/revel"
	"github.com/voyageth/frog/server/app/models"
	"log"
)

type CompanyController struct {
	AdminController
}

func (c CompanyController) Index() revel.Result {
	companyList, err := c.Txn.Select(models.Company{}, "select * from Company limit 10")
	log.Println(companyList)
	log.Println(err)
	if err != nil {
		log.Println(err.Error())
		c.Flash.Error(c.Message("server.error"))
		c.Validation.Keep()
		c.FlashParams()
	}
	return c.Render()
}