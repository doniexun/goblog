package admin

import (
	"github.com/astaxie/beego"
)

// BaseController 后台 controller 的父类
type BaseController struct {
	beego.Controller

	userid   int    // 登录用户的 id
	username string // 登录用户的 名称
}

// Get 一个例子
func (c *BaseController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["USER_AGENT"] = c.Ctx.Input.Header("user-agent")
	c.TplName = "index.tpl"
}
