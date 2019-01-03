package front

import (
	"strconv"

	"github.com/doniexun/goblog/models"
)

// MainController 前台的 controller
type MainController struct {
	BaseController
}

// Index 列出指定数目的文章
func (c *MainController) Index() {
	var (
		list     []*models.Post
		err      error
		page     int
		pagesize int
	)

	if page, err = strconv.Atoi(c.Ctx.Input.Param(":page")); err != nil || page < 1 {
		page = 1
	}

	// assume value of pagesize is 10
	// load this value from cache or db. [TODO]
	pagesize = 10

	query := new(models.Post).Query()
	count, _ := query.Count()
	if count > 0 {
		query.Limit(pagesize, (page-1)*pagesize).All(&list)
	}

	c.Data["count"] = count
	c.Data["list"] = list
	c.TplName = "index.tpl" // 相对路径：相对于工程根目录 views/
}
