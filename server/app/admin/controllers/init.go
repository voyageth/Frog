package controllers

import (
	"frog/server/app/controllers"
	"github.com/revel/revel"
)

func init() {
	revel.OnAppStart(controllers.InitDB)
	revel.InterceptMethod((*controllers.GorpController).Begin, revel.BEFORE)
	revel.InterceptMethod((*controllers.GorpController).Commit, revel.AFTER)
	//revel.InterceptMethod((*GorpController).Rollback, revel.FINALLY)
}
