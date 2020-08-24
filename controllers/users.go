package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"loveHome2020/models"
	"os"
	"path"
	"strconv"
	"time"
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

	//插入数据库
	o := orm.NewOrm()
	user := models.User{}
	user.Password_hash = resp["password"].(string)
	user.Name = resp["mobile"].(string)
	user.Mobile = resp["mobile"].(string)
	id, err := o.Insert(&user)
	if err != nil {
		resp["errno"] = 4002
		resp["errmsg"] = "注册失败"
	}
	beego.Info("reg success, id =", id)
	resp["errno"] = 0
	resp["errmsg"] = "注册成功"

	this.SetSession("name", user.Name)
	this.SetSession("user_id", user.Id)
	this.SetSession("mobile", user.Mobile)

	this.RetData(resp)
}

func (this *UserController) GetUserData() {
	resp := make(map[string]interface{})
	defer this.RetData(resp)

	// 从Session中获取用户ID
	user_id := this.GetSession("user_id")

	userModel := models.User{}
	// 查找用户信息
	oneErr := orm.NewOrm().QueryTable("user").Filter("id", user_id).One(&userModel, "id", "name", "mobile","real_name", "id_card", "avatar_url")

	if oneErr != nil {
		resp["errno"] = models.RECODE_NODATA
		resp["errmsg"] = models.RecodeText(models.RECODE_NODATA)
		return
	}

	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	resp["data"] = userModel

	return
}

// 上传图片
func (this *UserController) PostAvatar() {
	resp := make(map[string]interface{})
	defer this.RetData(resp)

	// 获取上传数据
	f, h, err := this.GetFile("avatar")
	if err != nil {
		resp["errno"] = models.RECODE_REQERR
		resp["errmsg"] = models.RecodeText(models.RECODE_REQERR)

		return
	}
	defer f.Close()

	// 获取文件后缀
	ext := path.Ext(h.Filename)
	// 当前秒级时间戳
	filetime := time.Now().Unix()
	filetimeString := strconv.FormatInt(filetime, 10)
	// 获取session中的userid
	userid := this.GetSession("user_id")
	useridString := strconv.Itoa(userid.(int))

	dir := "static/upload/" + useridString + "/"

	// 创建目录 os.ModePerm: 0777权限
	os.MkdirAll(dir, os.ModePerm)

	// 图片名称
	filename := dir + filetimeString + ext

	// 保存图片
	SaveToFileErr := this.SaveToFile("avatar", filename)
	if SaveToFileErr != nil {
		resp["errno"] = 111
		resp["errmsg"] = "图片保存失败"

		return
	}
	upNum, upErr := orm.NewOrm().QueryTable("user").Filter("id", userid).Update(orm.Params{"avatar_url": filename})
	if upErr != nil || upNum == 0 {
		resp["errno"] = models.RECODE_USERERR
		resp["errmsg"] = models.RecodeText(models.RECODE_USERERR)

		return
	}

	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	resp["data"] = filename

	return
}
