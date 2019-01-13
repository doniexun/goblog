package admin

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/doniexun/goblog/models"
	"github.com/lisijie/goblog/util"
)

// BaseController 后台 controller 的父类
type BaseController struct {
	beego.Controller

	userID         int64  // 登录用户的 id
	userName       string // 登录用户的 名称
	controllerName string // 控制器名称
	actionName     string // 操作名称
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
	c.controllerName, c.actionName = c.GetControllerAndAction()
	c.auth()
}

// auth 通过 cookie 鉴权
func (c *BaseController) auth() {
	authstr := c.Ctx.GetCookie("auth")
	authcookie := strings.Split(authstr, "|")
	if len(authcookie) == 2 {
		idstr, password := authcookie[0], authcookie[1]
		if userid, err := strconv.Atoi(idstr); err == nil {
			if userid > 0 {
				user := &models.User{}
				user.ID = int64(userid)
				if user.Read() == nil && password == util.Md5([]byte(user.Password)) {
					c.userID = int64(userid)
					c.userName = user.UserName
					log.Println("用户授权通过，当前用户 ID 是：" + idstr)
				}
			}
		}
	}

	// 只有在登录页面或登出页面，允许用户编号为0
	// 其他页面，只要未能获取到用户编号，就应该直接跳转
	if c.userID == 0 && c.controllerName != "AccountController" {
		log.Println("用户未登录，请先登录！")
		c.BackToClientReponse(false, "用户未登录，请先登录！")
		c.Redirect("/admin/account/login", 302)
	}

	// [TODO] 为了添加demo用户，在此处放行 actionName == AddDemoUser，正式版本应删掉
	if c.userID == 0 && (c.controllerName == "AccountController" && c.actionName != "Login" && c.actionName != "Logout" && c.actionName != "AddDemoUser") {
		log.Println("用户未登录，请先登录！")
		c.BackToClientReponse(false, "用户未登录，请先登录！")
		c.Redirect("/admin/account/login", 302)
	}
}

// ClientIP 获取连接客户端 IP 地址
func (c *BaseController) ClientIP() string {
	// [TODO] 获取到的 IP 地址为 [::1]:36064
	addr := c.Ctx.Request.RemoteAddr
	addrArray := strings.Split(addr, ":")
	return addrArray[0]
}

// BackToClientReponse 返回给客户端 json 格式的响应信息
func (c *BaseController) BackToClientReponse(status bool, msg string) {
	jsonstr := make(map[string]interface{})
	jsonstr["status"] = status
	jsonstr["msg"] = msg
	if data, err := json.Marshal(jsonstr); err != nil {
		log.Println(err.Error())
		c.Ctx.WriteString(err.Error())
	} else {
		c.Ctx.WriteString(string(data))
	}
}

// BackToClientData 返回给客户端 json 格式的数据
func (c *BaseController) BackToClientData(instance interface{}) {
	data, err := json.Marshal(instance)
	if err != nil {
		log.Println(err.Error())
		c.Ctx.WriteString(err.Error())
	}
	c.Ctx.WriteString(string(data))
}
