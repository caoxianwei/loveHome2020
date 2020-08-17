package controllers

import (
	"github.com/astaxie/beego"
	"loveHome2020/models"
)

type SessionController struct {
	beego.Controller
}

func (this *SessionController) RetData(resp map[string]interface{}) {
	this.Data["json"] = &resp
	this.ServeJSON()
}

func (this *SessionController) SessionData() {
	beego.Info("GetHouseIndex function")

	resp := make(map[string]interface{})

	user :=models.User{}
	user.Name = "eeeewww1"

	resp["errno"] = 0
	resp["errmsg"] = "OK"
	resp["data"] = user
	this.RetData(resp)
}