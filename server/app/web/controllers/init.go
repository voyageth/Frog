package controllers

import (
	"github.com/revel/revel"
	"frog/server/app/controllers"
)

func init() {
	revel.OnAppStart(controllers.InitDB)
	revel.InterceptMethod((*controllers.GorpController).Begin, revel.BEFORE)
	revel.InterceptMethod((*controllers.GorpController).Commit, revel.AFTER)
	//revel.InterceptMethod((*GorpController).Rollback, revel.FINALLY)
}
