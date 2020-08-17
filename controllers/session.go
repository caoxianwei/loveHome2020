package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
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
	beego.Info("GetSessionData function")
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

func (this *SessionController) DeleteSession() {
	beego.Info("DeleteSession function")

	resp := make(map[string]interface{})
	defer this.RetData(resp)
	this.DelSession("name")

	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
}

func (this *SessionController) Login() {
	beego.Info("Login function")
// 1.得到用户信息
// 2.判断信息是否合法
// 3.合数据库匹配账号密码是否正确
// 4.添加session
// 5.返回用户信息

	resp := make(map[string]interface{})
	defer this.RetData(resp)

	userModel := models.User{}
	//获取前端传过来的json数据
	jsonErr := json.Unmarshal(this.Ctx.Input.RequestBody, &userModel)
	if jsonErr != nil {
		resp["errno"] = models.RECODE_NODATA
		resp["errmsg"] = models.RecodeText(models.RECODE_NODATA)
		//resp["data"] = userModel
		return
	}
	queryErr := orm.NewOrm().QueryTable("user").Filter("mobile", userModel.Mobile).Filter("password_hash", userModel.Password_hash).One(&userModel)
	if queryErr == orm.ErrNoRows {
		resp["errno"] = models.RECODE_NODATA
		resp["errmsg"] = models.RecodeText(models.RECODE_NODATA)
		return
	}
	// 记录Session
	this.SetSession("name", userModel.Name)
	this.SetSession("user_id", userModel.Id)
	this.SetSession("mobile", userModel.Mobile)

	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
}