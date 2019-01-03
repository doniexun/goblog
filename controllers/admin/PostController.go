package admin

import (
	"github.com/astaxie/beego"
)

// PostController 用户账户 controller
type PostController struct {
	beego.Controller
}

// Get 一个例子
func (c *PostController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["USER_AGENT"] = c.Ctx.Input.Header("user-agent")
	c.TplName = "index.tpl"
}

// List 指定页的文章列表
func (c *PostController) List() {

}

// Add 添加文章
func (c *PostController) Add() {

}

// Delete 删除文章
func (c *PostController) Delete() {

}

// Update 修改文章
func (c *PostController) Update() {

}
