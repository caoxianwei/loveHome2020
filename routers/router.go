package routers

import (
	"loveHome2020/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/api/v1.0/areas", &controllers.AreaController{}, "get:GetArea")
	// Hourse Controller
	beego.Router("/api/v1.0/houses/index", &controllers.HouseIndexController{}, "get:GetHouseIndex")
	beego.Router("/api/v1.0/user/houses", &controllers.HouseIndexController{}, "get:GetHouseData")
	beego.Router("/api/v1.0/houses", &controllers.HouseIndexController{}, "post:PostHouseData")
	beego.Router("/api/v1.0/houses/:id([0-9]+)", &controllers.HouseIndexController{}, "get:HouseData")

	// Session Controller
	beego.Router("/api/v1.0/session", &controllers.SessionController{}, "get:SessionData;delete:DeleteSession")
	// User Controller
	beego.Router("/api/v1.0/users", &controllers.UserController{}, "post:Reg")
	beego.Router("/api/v1.0/user", &controllers.UserController{}, "get:GetUserData")
	beego.Router("api/v1.0/user/name", &controllers.UserController{}, "put:UpUserName")
	beego.Router("/api/v1.0/user/auth", &controllers.UserController{}, "get:GetUserData;post:UserAuth")
	beego.Router("/api/v1.0/sessions", &controllers.SessionController{}, "post:Login")

	beego.Router("/api/v1.0/user/avatar", &controllers.UserController{}, "post:PostAvatar")

	//Orders
	beego.Router("/api/v1.0/orders", &controllers.OrdersController{}, "post:PostOrders")
	beego.Router("/api/v1.0/user/orders", &controllers.OrdersController{}, "get:GetOrdersData")
}
