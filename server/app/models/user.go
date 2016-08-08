package models

import (
	"fmt"
	"github.com/revel/revel"
	"regexp"
	"log"
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

var userRegex = regexp.MustCompile(`\w[-._\w]*\w@\w[-._\w]*\w\.\w{2,3}`)

func (user *User) Validate(c *revel.Controller, v *revel.Validation) {
	log.Print("user : " + user.Id)
	v.Required(user.Id).Message(c.Message("user.register.email.required"))
	v.MinSize(user.Id, 3).Message(c.Message("user.register.email.tooShort"))
	v.MaxSize(user.Id, 30).Message(c.Message("user.register.email.tooLong"))
	v.Match(user.Id, userRegex).Message(c.Message("user.register.email.format"))

	v.Required(user.Name).Message(c.Message("user.register.name.required"))
	v.MinSize(user.Name, 3).Message(c.Message("user.register.name.tooShort"))
	v.MaxSize(user.Name, 20).Message(c.Message("user.register.name.tooLong"))

	ValidatePassword(c, v, user.Password)
}

func ValidatePassword(c *revel.Controller, v *revel.Validation, password string) {
	v.Required(password).Message(c.Message("user.register.password.required"))
	v.MinSize(password, 4).Message(c.Message("user.register.password.tooShort"))
	v.MaxSize(password, 20).Message(c.Message("user.register.password.tooLong"))
}
