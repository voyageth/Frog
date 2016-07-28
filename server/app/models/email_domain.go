package models

import (
	"fmt"
	"github.com/go-gorp/gorp"
	"github.com/revel/revel"
)

type EmailDomain struct {
	No        int
	Domain    string
	CompanyNo int

	// Transient
	Company *Company
}

func (emailDomain *EmailDomain) Validate(v *revel.Validation) {
	v.Required(emailDomain.Company)

	v.Check(
		emailDomain.Domain,
		revel.Required{},
		revel.MinSize{1},
		revel.MaxSize{200},
	)
}

func (emailDomain *EmailDomain) String() string {
	return fmt.Sprintf("EmailDomain(Domain : %s, Company : %s)", emailDomain.Domain, emailDomain.Company)
}

func (emailDomain *EmailDomain) PreInsert(_ gorp.SqlExecutor) error {
	emailDomain.CompanyNo = emailDomain.Company.No
	return nil
}

func (emailDomain *EmailDomain) PostGet(exe gorp.SqlExecutor) error {
	var (
		obj interface{}
		err error
	)

	obj, err = exe.Get(Company{}, emailDomain.No)
	if err != nil {
		return fmt.Errorf("Error loading a emailDomain's company (%d): %s", emailDomain.CompanyNo, err)
	}
	emailDomain.Company = obj.(*Company)

	return nil
}
