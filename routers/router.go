package routers

import (
	"github.com/astaxie/beego"
	"github.com/doniexun/goblog/controllers/admin"
	"github.com/doniexun/goblog/controllers/front"
)

func init() {
	beego.Router("/", &front.MainController{}, "get:Index")
	beego.Router("/post/:id:int", &front.MainController{}, "get:Show")

	beego.Router("/admin/account/login", &admin.AccountController{}, "post:Login")
	beego.Router("/admin/account/logout", &admin.AccountController{}, "get:Logout")
}
