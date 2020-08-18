package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"loveHome2020/models"
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

	areasModel := []models.Area{}

	//errCache := cache_conn.Put("aaa", "bbbbaaa", time.Second * 3600)
	//if errCache != nil {
	//	beego.Error("cacha err")
	//}
	//beego.Info("cache_conn.aa = ", cache_conn.Get("aaa"))
	//fmt.Printf("cache_conn[aa] = %s\n ",cache_conn.Get("aaa"))

	// 从redis中取数据
	if cache_conn.IsExist("area") {
		if redisArea := cache_conn.Get("area"); redisArea != nil {
			resp["errno"] = models.RECODE_OK
			resp["errmsg"] = models.RecodeText(models.RECODE_OK)
			//area := json.Unmarshal(redisArea, &models.Area{})
			json.Unmarshal(redisArea.([]byte), &areasModel)
			resp["data"] = areasModel

			beego.Info("取到数据了")

			return
		}
	}

	// 从数据库中读取数据
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

	// 将area数据存入Redis
	jsonStr, jsonErr := json.Marshal(areasModel)
	if jsonErr != nil {
		resp["error"] = models.RECODE_DATAERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DATAERR)
		return
	}
	redisPutErr := cache_conn.Put("area", jsonStr, time.Second * 3600)
	if redisPutErr != nil {
		resp["errno"] = 1234
		resp["errmsg"] = "缓存操作失败了"
		return
	}

	beego.Info("数据查询成功, resp=", resp, "num =", num)
	//json_str := json.Marshal(resp)
	//c.Ctx.WriteString(json_str)
	//c.Data["json"] = resp
	//c.ServeJSON()
}
