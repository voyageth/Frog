package admin_controllers

import (
	"github.com/revel/revel"
	"frog/server/app/controllers/web"
	"log"
)

type AdminController struct {
	web_controllers.WebController
}

func AdminUserCheckFilter(c *revel.Controller, fc []revel.Filter) {
	// TODO impl
	log.Println("AdminUserCheckFilter!")
	fc[0](c, fc[1:])
}