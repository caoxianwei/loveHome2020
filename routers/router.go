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

	// Session Controller
	beego.Router("/api/v1.0/session", &controllers.SessionController{}, "get:SessionData;delete:DeleteSession")
	// User Controller
	beego.Router("/api/v1.0/users", &controllers.UserController{}, "post:Reg")
	beego.Router("/api/v1.0/sessions", &controllers.SessionController{}, "post:Login")
}
