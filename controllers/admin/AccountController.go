package admin

import (
	"fmt"
	"log"
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
	cookie := strconv.FormatInt(user.ID, 10) + "|" + authkey
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

// ExistByID 判断用户是否存在（通过用户 ID 来判断）
func (c *AccountController) ExistByID(userid int64) bool {
	user := models.User{ID: userid}
	if user.Read() != nil {
		return false
	}
	return true
}

// ExistByName 判断用户是否存在（通过用户名来判断）
func (c *AccountController) ExistByName(userName string) bool {
	user := models.User{UserName: userName}
	if err := user.Read("UserName"); err != nil {
		return false
	}
	return true
}

// AddDemoUser 添加测试用户
func (c *AccountController) AddDemoUser() {

	demoName := "windness"
	if c.ExistByName(demoName) {
		log.Println("用户添加失败，用户 " + demoName + " 已经存在")
		c.BackToClientReponse(false, "用户添加失败，用户 "+demoName+" 已经存在")
		return
	}

	var user models.User
	user.UserName = demoName
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
	if err := user.Insert(); err != nil {
		log.Println(err.Error())
		c.BackToClientReponse(false, "用户添加失败"+err.Error())
		return
	}

	c.BackToClientReponse(true, "用户"+user.UserName+"添加成功")

}
