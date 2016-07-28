package models

import (
	"fmt"
	"github.com/revel/revel"
	"regexp"
)

type User struct {
	No             int
	Id             string // email
	Name           string
	Password       string
	HashedPassword []byte
	EmailDomain    *EmailDomain
}

func (user *User) String() string {
	return fmt.Sprintf("User(id : %s, name : %s)", user.Id, user.Name)
}

var userRegex = regexp.MustCompile("^\\w*$")

func (user *User) Validate(v *revel.Validation) {
	v.Check(
		user.Id,
		revel.Required{},
		revel.MaxSize{15},
		revel.MinSize{4},
		revel.Match{userRegex},
	)

	v.Check(
		user.Name,
		revel.Required{},
		revel.MaxSize{100},
	)

	ValidatePassword(v, user.Password).Key("user.Password")
}

func ValidatePassword(v *revel.Validation, password string) *revel.ValidationResult {
	return v.Check(
		password,
		revel.Required{},
		revel.MaxSize{15},
		revel.MinSize{5},
	)
}
