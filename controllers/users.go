package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) RetData(resp map[string]interface{}) {
	this.Data["json"] = &resp
	this.ServeJSON()
}

func (this *UserController) Reg() {
	beego.Info("Reg function")

	resp := make(map[string]interface{})
	json.Unmarshal(this.Ctx.Input.RequestBody, &resp)

	beego.Info(`resp["mobile"] =`, resp["mobile"])
	beego.Info(`resp["password"] =`, resp["password"])
	beego.Info(`resp["sms_code"] =`, resp["sms_code"])

	this.RetData(resp)
}