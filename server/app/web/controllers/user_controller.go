package controllers

import (
	"github.com/revel/revel"
	"log"
	"frog/server/app/models"
	"golang.org/x/crypto/bcrypt"
	"github.com/go-sql-driver/mysql"
	"strings"
	"frog/server/app/controllers"
)

type UserController struct {
	controllers.App
}

const SESSION_KEY_LOGIN = "frog_login"

func (c UserController) Index() revel.Result {
	return c.Render()
}

func (c UserController) LoginRequest(userEmail string, password string) revel.Result {
	var user models.User
	err := c.Txn.SelectOne(&user, "select * from User where Email=?", userEmail)
	log.Println(user)
	log.Println(err)
	if err != nil {
		log.Println(err.Error())
		c.Flash.Error(c.Message("server.error"))
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(UserController.Index)
	}

	hashedPassword := user.HashedPassword
	passwordCompareErr := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if passwordCompareErr == nil {
		// TODO #6 more complicate login cookie
		c.Session[SESSION_KEY_LOGIN] = userEmail
		return c.Redirect(controllers.App.Index)
	} else {
		log.Println(passwordCompareErr.Error())
		c.Flash.Error(c.Message("user.login.password.wrong"))
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(UserController.Index)
	}
}

func (c UserController) Logout() revel.Result {
	delete(c.Session, SESSION_KEY_LOGIN)
	return c.Redirect(controllers.App.Index)
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
		log.Println(err.Error())
		c.Flash.Error(c.Message("server.error"))
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(UserController.Register)
	}
	user.HashedPassword = hashedPassword

	err = c.Txn.Insert(user)
	if err != nil {
		log.Println(err.Error())
		if mysqlError, ok := err.(*mysql.MySQLError); ok {
			if mysqlError.Number == 1062 {
				if strings.Contains(err.Error(), "for key 'Email'") {
					c.Flash.Error(c.Message("user.register.email.duplicated"))
					c.Validation.Keep()
					c.FlashParams()
					return c.Redirect(UserController.Register)
				}
				if strings.Contains(err.Error(), "for key 'Name'") {
					c.Flash.Error(c.Message("user.register.name.duplicated"))
					c.Validation.Keep()
					c.FlashParams()
					return c.Redirect(UserController.Register)
				}
			}
		}
		c.Flash.Error(c.Message("server.error"))
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(UserController.Register)
	}

	// TODO #5 request email verification

	return c.Redirect(UserController.Index, user.Name)
}
