package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	_ "loveHome2020/models"
	_ "loveHome2020/routers"
	"net/http"
	"strings"
)

func main() {
	ignoreStaticPath()
	// 第一种方法开启Session
	beego.BConfig.WebConfig.Session.SessionOn = true

	beego.Run()
}

func ignoreStaticPath() {
	// 透明static
	beego.InsertFilter("/", beego.BeforeRouter, TransparentStatic)
	beego.InsertFilter("/*", beego.BeforeRouter, TransparentStatic)
}

func TransparentStatic(ctx *context.Context) {
	orpath := ctx.Request.URL.Path
	beego.Debug("request url: ", orpath)
	// 如果请求uri还有api字段，说明是指令应该取消静态资源路径重定向
	if strings.Index(orpath, "api") >= 0 {
		return
	}

	http.ServeFile(ctx.ResponseWriter, ctx.Request, "static/html"+orpath)
}
