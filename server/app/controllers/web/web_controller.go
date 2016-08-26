package web_controllers

import (
	"github.com/revel/revel"
	"github.com/voyageth/frog/server/app/controllers"
	"log"
)

type WebController struct {
	controllers.App
}

func UserLoginFilter(c *revel.Controller, fc []revel.Filter) {
	userEmail := c.Session[SESSION_KEY_LOGIN]
	log.Println("UserLoginFilter. userEmail : " + userEmail)
	c.RenderArgs["userEmail"] = userEmail

	fc[0](c, fc[1:])
}