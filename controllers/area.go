package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"loveHome2020/models"
	_ "github.com/go-sql-driver/mysql"
)

type AreaController struct {
	beego.Controller
}

func (this *AreaController) RetData(resp map[string]interface{}) {
	this.Data["json"] = &resp
	this.ServeJSON()
}

func (c *AreaController) GetArea() {
	beego.Info("GetArea sucusse")

	var resp map[string]interface{}
	defer c.RetData(resp)

	// 从数据库中读取数据
	areasModel := []models.Area{}

	o := orm.NewOrm()
	_, allErr := o.QueryTable("area").All(&areasModel)
	if allErr != nil {
		beego.Info("数据错误")
		resp["errno"] = 4001
		resp["errmsg"] = "查询失败"

		return
	}
	if allErr != nil {
		resp["errno"] = 4002
		resp["errmsg"] = "没有查到数据"
		return
	}
	resp["errno"] = 0
	resp["errmsg"] = "OK"
	resp["data"] = areasModel

	beego.Info("数据查询成功")
	//json_str := json.Marshal(resp)
	//c.Ctx.WriteString(json_str)
	//c.Data["json"] = resp
	//c.ServeJSON()
}
