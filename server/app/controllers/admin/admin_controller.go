package admin_controllers

import (
	"github.com/revel/revel"
	"github.com/voyageth/frog/server/app/controllers/web"
	"log"
)

type AdminController struct {
	web_controllers.WebController
}

// TODO read from properties?
const ADMIN_USER = "frog_login"

func AdminUserCheckFilter(c *revel.Controller, fc []revel.Filter) {
	log.Println("AdminUserCheckFilter!")

	userEmail := c.Session[web_controllers.SESSION_KEY_LOGIN]
	if (ADMIN_USER == userEmail) {
		c.RenderArgs["isAdminUser"] = true
	}

	fc[0](c, fc[1:])
}