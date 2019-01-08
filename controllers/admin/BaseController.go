package admin

import (
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/doniexun/goblog/models"
	"github.com/lisijie/goblog/util"
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

// Prepare 在此完成鉴权操作，所有后台操作之前都会先调用此函数
func (c *BaseController) Prepare() {
	c.auth()
}

// auth 通过 cookie 鉴权
func (c *BaseController) auth() {
	authcookie := strings.Split(c.Ctx.GetCookie("auth"), "|")
	if len(authcookie) == 2 {
		idstr, password := authcookie[0], authcookie[1]
		userid, _ := strconv.Atoi(idstr)
		if userid > 0 {
			var user models.User
			user.ID = userid
			if user.Read() == nil && password == util.Md5([]byte(user.Password)) {
				c.userid = userid
				c.username = user.UserName
			}
		}
	}
}
