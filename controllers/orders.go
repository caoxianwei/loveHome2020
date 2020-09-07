package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"loveHome2020/models"
)

type OrdersController struct {
	beego.Controller
}

func (this *OrdersController) RetData(resp map[string]interface{}) {
	this.Data["json"] = &resp
	this.ServeJSON()
}
func (this *OrdersController) GetOrdersData() {
	beego.Info("GetOrdersData function")

	resp := make(map[string]interface{})
	defer this.RetData(resp)

	user_id := this.GetSession("user_id")
	ordersModel := []models.OrderHouse{}
	ormer := orm.NewOrm()
	_, err := ormer.QueryTable("order_house").Filter("User", user_id).RelatedSel().All(&ordersModel)
	if err != nil {
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)

		return
	}

	role := this.GetString("role")
	if role == "custom" {
		orders := []models.House{}
		//order := models.OrderHouse{}

		o := orm.NewOrm()
		qs := o.QueryTable("OrderHouse")

		user := models.User{Id: user_id.(int)}
		qs.Filter("user__id", user_id.(int)).All(&orders)
		for _, order := range orders {
			order.User = &user
			o.LoadRelated(order, "User")
		}

		//order.User
		//o.LoadRelated(&order, "User")
	}
	respData := make(map[string]interface{})
	respData["order"] = ordersModel
	resp["data"] = respData

	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
}

func (this *OrdersController) PostOrders() {
	beego.Info("PostOrders function")

	resp := make(map[string]interface{})
	defer this.RetData(resp)

	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
}
