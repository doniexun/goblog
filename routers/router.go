package routers

import (
	"github.com/doniexun/goblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.BaseController{})
}
