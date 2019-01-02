package front

import (
	"strconv"
	
)

type MainController struct {
	BaseController
}

func (c *MainController) Index() {
	var (
		list []*models.Post
		err error
		page int
		pagesize int
	)
	
	if page, err = strconv.Atoi(c.Ctx.Input.Param(":page")); err != nil || page < 1 {
	page = 1
	}
	
		

}

