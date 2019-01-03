package routers

import (
	"github.com/astaxie/beego"
	"github.com/doniexun/goblog/controllers/admin"
	"github.com/doniexun/goblog/controllers/front"
)

func init() {
	beego.Router("/", &front.MainController{}, "get:Index")
	beego.Router("/post/:id:int", &front.MainController{}, "get:Show")

	beego.Router("/admin", &admin.AccountController{}, "get:Login")
}
