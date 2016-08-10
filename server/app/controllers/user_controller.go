package controllers

import (
	"github.com/revel/revel"
	"log"
	"frog/server/app/models"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	App
}

const SESSION_KEY_LOGIN = "frog_login"

func (c UserController) Index() revel.Result {
	return c.Render()
}

func (c UserController) Login(userEmail string, password string) revel.Result {
	c.Session[SESSION_KEY_LOGIN] = userEmail
	return c.Redirect(App.Index)
}

func (c UserController) Logout() revel.Result {
	delete(c.Session, SESSION_KEY_LOGIN)
	return c.Redirect(App.Index)
}

func (c UserController) Register() revel.Result {
	return c.Render()
}

func (c UserController) RegisterRequest(user *models.User) revel.Result {
	log.Print(revel.MessageLanguages())
	log.Print("user : ", user)

	user.Validate(c.Controller, c.Validation)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(UserController.Register)
	}

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	user.HashedPassword = hashedPassword

	err = c.Txn.Insert(user)
	if err != nil {
		c.Flash.Error(err.Error())
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(UserController.Register)
	}

	// TODO #5 request email verification

	return c.Redirect(UserController.Login, user.Name)
}
