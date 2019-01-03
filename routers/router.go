package routers

import (
	"github.com/astaxie/beego"
	"github.com/doniexun/goblog/controllers/front"
)

func init() {
	beego.Router("/", &front.MainController{}, "get:Index")
	beego.Router("/post/:id:int", &front.MainController{}, "get:Show")
}
