package main

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/doniexun/goblog/models"
	_ "github.com/doniexun/goblog/routers"
	"github.com/doniexun/goblog/util"
)

func main() {

	var post models.Post
	post.AuthorID = 1
	post.AuthorName = "windness"
	post.Title = "Test Post Title"
	post.Content = "A testing post's content"

	post.Insert()

	var user models.User
	user.UserName = "windness"
	user.Email = "windnessr@163.com"
	user.Password = util.Md5([]byte("123456"))
	user.RegisterTime = time.Now()
	user.LastLoginTime = time.Now()
	user.Insert()

	beego.Run()
}
