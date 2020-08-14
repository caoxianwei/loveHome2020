package controllers

import "github.com/astaxie/beego"

type AreaController struct {
	beego.Controller
}

func (c *AreaController) GetArea() {
	beego.Info("GetArea sucusse")
}
