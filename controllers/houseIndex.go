package controllers

import (
	"github.com/astaxie/beego"
	"loveHome2020/models"
)

type HouseIndexController struct {
	beego.Controller
}

func (this *HouseIndexController) RetData(resp map[string]interface{}) {
	this.Data["json"] = &resp
	this.ServeJSON()
}

func (this *HouseIndexController) GetHouseIndex() {
	beego.Info("GetHouseIndex function")

	resp := make(map[string]interface{})
	defer this.RetData(resp)

	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
}