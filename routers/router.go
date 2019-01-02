package routers

import (
	"github.com/doniexun/goblog/controllers/front"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &front.BaseController{})
}
