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
	defer this.RetData(resp)

	user :=models.User{}
	//user.Name = "eeeewww1"

	resp["errno"] = models.RECODE_DATAERR
	resp["errmsg"] = models.RecodeText(models.RECODE_DATAERR)
	name := this.GetSession("name")
	if name != nil {
		user.Name = name.(string)
		resp["errno"] = models.RECODE_OK
		resp["errmsg"] = models.RecodeText(models.RECODE_OK)
		resp["data"] = user
	}

}