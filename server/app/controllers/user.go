package controllers

import (
	"github.com/revel/revel"
	"log"
)

type User struct {
	App
}

const SESSION_KEY_LOGIN = "frog_login"

func (c User) Index() revel.Result {
	return c.Render()
}

func (c User) Login(userEmail string, password string) revel.Result {
	c.Session[SESSION_KEY_LOGIN] = userEmail
	return c.Redirect(App.Index)
}

func (c User) Logout() revel.Result {
	delete(c.Session, SESSION_KEY_LOGIN)
	return c.Redirect(App.Index)
}

func (c User) Register() revel.Result {
	return c.Render()
}

func (c User) RegisterRequest(userEmail string, userName string, password string) revel.Result {
	log.Print(revel.MessageLanguages())
	log.Print("userEmail : ", userEmail, ", userName : ", userName)
	c.Validation.Required(userEmail).Message(c.Message("user.register.email.required"))
	c.Validation.MinSize(userEmail, 3).Message(c.Message("user.register.email.tooShort"))
	c.Validation.MaxSize(userEmail, 30).Message(c.Message("user.register.email.tooLong"))
	// TODO email format check

	c.Validation.Required(userName).Message(c.Message("user.register.name.required"))
	c.Validation.MinSize(userName, 3).Message(c.Message("user.register.name.tooShort"))
	c.Validation.MaxSize(userName, 20).Message(c.Message("user.register.name.tooLong"))

	c.Validation.Required(password).Message(c.Message("user.register.password.required"))
	c.Validation.MinSize(password, 4).Message(c.Message("user.register.password.tooShort"))
	c.Validation.MaxSize(password, 20).Message(c.Message("user.register.password.tooLong"))

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(User.Register)
	}

	// TODO #4 store to db

	// TODO #5 request email verification

	return c.Redirect(User.Login, userName)
}
