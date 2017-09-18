package routers

import (
	"MobileMonitor/controllers"
	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/",&controllers.LoginController{},"get:LoginHtml")

	beego.AddNamespace(
		beego.NewNamespace("/v1",

			//登录相关
			beego.NSRouter("/login",&controllers.LoginController{},"get:LoginHtml"),
			beego.NSRouter("/loginApi",&controllers.LoginController{},"post:LoginApi"),

			//注册相关
			beego.NSRouter("/register",&controllers.RegController{},"get:RegHtml"),
			beego.NSRouter("/registerApi",&controllers.RegController{},"post:RegApi"),

			//主界面相关
			beego.NSNamespace("/main",
				//首页
				beego.NSRouter("/index",&controllers.MainController{},"get:Index"),
			),
		),
	)
}
