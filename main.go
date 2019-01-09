package main

import (
	"github.com/astaxie/beego"
	_ "github.com/doniexun/goblog/routers"
)

func main() {
	beego.Run()
}
