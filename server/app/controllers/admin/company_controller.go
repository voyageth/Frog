package admin_controllers

import (
	"github.com/revel/revel"
	"github.com/voyageth/frog/server/app/models"
	"log"
	"github.com/go-sql-driver/mysql"
	"strings"
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

func (c CompanyController) Register() revel.Result {
	return c.Render()
}

func (c CompanyController) RegisterRequest(company *models.Company) revel.Result {
	log.Print("company : ", company)

	company.Validate(c.Validation)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(CompanyController.Register)
	}

	err := c.Txn.Insert(company)
	if err != nil {
		log.Println(err.Error())
		if mysqlError, ok := err.(*mysql.MySQLError); ok {
			if mysqlError.Number == 1062 {
				if strings.Contains(err.Error(), "for key 'Name'") {
					c.Flash.Error(c.Message("company.register.name.duplicated"))
					c.Validation.Keep()
					c.FlashParams()
					return c.Redirect(CompanyController.Register)
				}
			}
		}
		c.Flash.Error(c.Message("server.error"))
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(CompanyController.Register)

	}

	return c.Redirect(CompanyController.Index)
}