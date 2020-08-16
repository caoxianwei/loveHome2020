package controllers

import "github.com/astaxie/beego"

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
	resp["errno"] = 4001
	resp["errmsg"] = "查询失败"
	this.RetData(resp)
}