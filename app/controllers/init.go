package controllers

import "github.com/revel/revel"

func init() {
	revel.InterceptMethod(App.CheckLogin, revel.BEFORE)
}
