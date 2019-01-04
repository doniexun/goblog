package admin

import (
	"fmt"
	"strconv"

	"github.com/doniexun/goblog/models"
	"github.com/lisijie/goblog/util"
)

// AccountController 用户账户 controller
type AccountController struct {
	BaseController
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

	fmt.Println(username)
	fmt.Println(password)

	if username != "" && password != "" {
		var user models.User
		user.Name = username
		if user.Read("name") != nil || user.Password != util.Md5([]byte(password)) {
			c.Data["errmsg"] = "账号或密码错误"
			fmt.Println("账号或密码错误")
		} else {
			authkey := util.Md5([]byte(user.Password))
			cookie := strconv.Itoa(user.Id) + "|" + authkey
			c.Ctx.SetCookie("auth", cookie)
			fmt.Println("账号和密码正确")
		}
	}

	// 跳转到后台主页 or 传回 json 字串[TODO]
	// ...
	c.Data["post"] = "test content"
	c.TplName = "admin/index.tpl"
}

// Logout 用户退出
func (c *AccountController) Logout() {
	c.Ctx.SetCookie("auth", "")

	// 跳转到登录页面 or 传回 json 字串[TODO]
	// ...
}

// Profile 用户信息
func (c *AccountController) Profile() {
	user := models.User{Id: c.userid}
	if err := user.Read(); err != nil {
		fmt.Println("The user not exists")
	}
	fmt.Println(user.Id)
	fmt.Println(user.Name)
	fmt.Println(user.Email)
	c.Data["user"] = user
	c.TplName = "admin/user.tpl"
}
