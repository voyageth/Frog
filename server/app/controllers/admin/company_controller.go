package admin_controllers

import (
	"github.com/revel/revel"
	"github.com/voyageth/frog/server/app/models"
	r "github.com/revel/revel"
	"github.com/go-sql-driver/mysql"
	"strings"
)

type CompanyController struct {
	AdminController
}

func (c CompanyController) Index() revel.Result {
	companyList, err := c.Txn.Select(models.Company{}, "select * from Company limit 10")
	r.TRACE.Println(companyList)
	r.TRACE.Println(err)
	if err != nil {
		r.WARN.Println(err.Error())
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
	r.TRACE.Print("company : ", company)

	company.Validate(c.Validation)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(CompanyController.Register)
	}

	err := c.Txn.Insert(company)
	if err != nil {
		r.WARN.Println(err.Error())
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