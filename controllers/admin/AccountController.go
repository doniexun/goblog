package admin

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/doniexun/goblog/models"
	"github.com/lisijie/goblog/util"
)

// AccountController 用户账户 controller
type AccountController struct {
	BaseController
}

// Register 用户注册
func (c *AccountController) Register() {
	username := c.GetString("username")
	password := c.GetString("password")
	email := c.GetString("email")

	// 对用户名、用户密码、email 进行合规性校验
	l := strings.Count(username, "") - 1
	if l < 2 || l > 16 {
		c.BackToClientReponse(false, "用户名长度限制在 2-16 个字符")
	}

	// 对密码、email进行校验
	// ...[TODO]

	// 对用户名是否已注册进行校验
	var user models.User
	user.UserName = username
	if user.Query().Filter("username", username).One(&user); user.ID > 0 {
		c.BackToClientReponse(false, "用户名 "+username+" 已被注册")
	}

	//if err := user.Read(); err != nil {
	//	c.Data["errmsg"] = "用户名已被注册"
	//	return //[TODO] 向客户端发送 json
	//}

	user.Email = email
	user.Password = util.Md5([]byte(password))
	if err := user.Insert(); err != nil {
		c.BackToClientReponse(false, "用户 "+username+" 注册失败；失败原因："+err.Error())
	}

	c.BackToClientReponse(true, "用户 "+username+" 注册成功")
}

// Login 用户登录
func (c *AccountController) Login() {

	username := c.GetString("username")
	password := c.GetString("password")

	fmt.Println(username)
	fmt.Println(password)

	if username == "" || password == "" {
		c.BackToClientReponse(false, "用户名和密码不得为空")
		return // [TODO] return or c.StopRun() ?
	}

	var user models.User
	user.UserName = username
	if err := user.Read("user_name"); err != nil { // "user_name" 或 "username" 均可以
		c.BackToClientReponse(false, "账号不存在")
		return // [TODO] return or c.StopRun() ?
	}

	if user.Password != util.Md5([]byte(password)) {
		c.BackToClientReponse(false, "密码错误")
		return // [TODO] return or c.StopRun() ?
	}

	authkey := util.Md5([]byte(user.Password))
	cookie := strconv.Itoa(user.ID) + "|" + authkey
	c.Ctx.SetCookie("auth", cookie)
	c.BackToClientReponse(true, "登录成功")
}

// Logout 用户退出
func (c *AccountController) Logout() {
	c.Ctx.SetCookie("auth", "")
	c.BackToClientReponse(true, "注销成功")
}

// Profile 用户信息
func (c *AccountController) Profile() {
	user := models.User{ID: c.userID}
	if err := user.Read(); err != nil {
		fmt.Println("The user not exists")
	}

	fmt.Println(user.ID)
	fmt.Println(user.UserName)
	fmt.Println(user.Email)
	c.Data["user"] = user
	c.TplName = "admin/user.tpl"
}

// AddDemoUser 添加测试用户
func (c *AccountController) AddDemoUser() {
	var user models.User
	user.UserName = "windness"
	user.NickName = "doniexun"
	user.Password = util.Md5([]byte("123456"))
	user.Email = "windnessr@163.com"
	user.QQ = "1758953369"
	user.Wechat = "splendidream"
	user.Cellphone = "13800138000"
	user.RegisterTime = time.Now()
	user.RegsiterIP = c.ClientIP()
	user.LastLoginTime = time.Now()
	user.LastLoginIP = c.ClientIP()
	user.Insert()
}
