package controllers

import (
	"github.com/revel/revel"
	"log"
)

func UserLoginFilter(c *revel.Controller, fc []revel.Filter) {
	userEmail := c.Session[SESSION_KEY_LOGIN]
	log.Println("UserLoginFilter. userEmail : " + userEmail)
	c.RenderArgs["userEmail"] = userEmail

	fc[0](c, fc[1:])
}