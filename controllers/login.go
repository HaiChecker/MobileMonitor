package controllers


type LoginController struct {
	BaseController
}

/**
 *	登录的页面
 */
func (login *LoginController) LoginHtml() {
	//TODO 等待写入登录界面
	login.TplName = "index.tpl"
}

/**
 *	登录的Api
 */
func (login *LoginController) LoginApi() {
	//TODO 等待写入登录Api
}