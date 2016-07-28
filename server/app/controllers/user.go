package controllers

import (
	"github.com/revel/revel"
	"log"
)

type User struct {
	*revel.Controller
}

func (c User) Index() revel.Result {
	return c.Render()
}

func (c User) Login(userEmail string, password string) revel.Result {
	return c.Redirect(App.Index)
}

func (c User) Logout() revel.Result {
	return c.Redirect(App.Index)
}

func (c User) Register() revel.Result {
	return c.Render()
}

func (c User) RegisterRequest(userEmail string, userName string, password string) revel.Result {
	log.Print("userEmail : ", userEmail, ", userName : ", userName)
	c.Validation.Required(userEmail).Message("Email is required!")
	c.Validation.MinSize(userEmail, 3).Message("Email is too short!")
	c.Validation.MaxSize(userEmail, 30).Message("Email is too long!")
	// TODO email format check

	c.Validation.Required(userName).Message("Name is required!")
	c.Validation.MinSize(userName, 3).Message("Name is too short!")
	c.Validation.MaxSize(userName, 20).Message("Name is too long!")

	c.Validation.Required(password).Message("Password is required!")
	c.Validation.MinSize(password, 4).Message("Password is too short!")
	c.Validation.MaxSize(password, 20).Message("Password is too long!")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(User.Register)
	}

	return c.Redirect(User.Login, userName)
}
