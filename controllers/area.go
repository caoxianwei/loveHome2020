package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/cache/redis"
	"loveHome2020/models"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type AreaController struct {
	beego.Controller
}

func (this *AreaController) RetData(resp map[string]interface{}) {
	this.Data["json"] = &resp
	this.ServeJSON()
}

func (this *AreaController) GetArea() {
	beego.Info("GetArea sucusse")

	//var resp map[string]interface{}
	resp := make(map[string]interface{})
	defer this.RetData(resp)

	// 从Redis中读取数据
	// 初始化一个全局变量对象
	cache_conn, err := cache.NewCache("redis", `{"key":"lovehome","conn":":6379","dbNum":"0"}`)
	if err != nil {
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}

	errCache := cache_conn.Put("aaa", "bbbbaaa", time.Second * 3600)
	if errCache != nil {
		beego.Error("cacha err")
	}
	beego.Info("cache_conn.aa = ", cache_conn.Get("aaa"))
	fmt.Printf("cache_conn[aa] = %s\n ",cache_conn.Get("aaa"))
	// 从数据库中读取数据
	areasModel := []models.Area{}
	o := orm.NewOrm()
	num, allErr := o.QueryTable("area").All(&areasModel)
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

	beego.Info("数据查询成功, resp=", resp, "num =", num)
	//json_str := json.Marshal(resp)
	//c.Ctx.WriteString(json_str)
	//c.Data["json"] = resp
	//c.ServeJSON()
}
