package front

import (
	"github.com/astaxie/beego"
)

// BaseController 所有前台 controller 的父类.
type BaseController struct {
	beego.Controller
}

// Get a example function.
func (c *BaseController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["USER_AGENT"] = c.Ctx.Input.Header("user-agent")
	c.TplName = "index.tpl"
}
