package admin

import (
	"github.com/astaxie/beego"
	"github.com/doniexun/goblog/models"
	"github.com/lisijie/goblog/util"
)

// AccountController 用户账户 controller
type AccountController struct {
	beego.Controller
}

// Get 一个例子
func (c *AccountController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["USER_AGENT"] = c.Ctx.Input.Header("user-agent")
	c.TplName = "index.tpl"
}

// Login 用户登录
func (c *AccountController) Login() {
	username := c.GetString("username")
	password := c.GetString("password")

	if username != "" && password != "" {
		var user models.User
		if user.Read("user_name") != nil || user.Password != util.Md5([]byte(password)) {
			c.Data["errmsg"] = "账号或密码错误"
		} else {
			authkey := util.Md5([]byte(user.Password))
			c.Ctx.SetCookie("auth", authkey)
		}
	}

	c.Data["post"] = "test content"
	c.TplName = "admin/index.tpl"
}

// Logout 用户退出
func (c *AccountController) Logout() {

}

// Profile 用户信息
func (c *AccountController) Profile() {

}
