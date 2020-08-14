package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"loveHome2020/models"
)

type AreaController struct {
	beego.Controller
}

func (c *AreaController) GetArea() {
	beego.Info("GetArea sucusse")

	resp := map[string]interface{}

	// 从数据库中读取数据
	area := models.Area{}

	o := orm.NewOrm()
	err := o.Read(&area)
	if err != nil {
		beego.Info("数据错误")
		resp["errno"] = 4001
		resp["errmsg"] = "查询失败"

		return
	}
	resp["errno"] = 0
	resp["errmsg"] = "OK"
	resp["data"] = &area

	json_str := json.Marshal(resp)
	c.Ctx.WriteString(json_str)
	c.Data["json"] = resp
	c.ServeJSON()
}
