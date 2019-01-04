package main

import (
	"github.com/astaxie/beego"
	"github.com/doniexun/goblog/models"
	_ "github.com/doniexun/goblog/routers"
	"github.com/doniexun/goblog/util"
)

func main() {

	var post models.Post
	post.AuthorId = 1
	post.AuthorName = "windness"
	post.Title = "Test Post Title"
	post.Content = "A testing post's content"

	post.Insert()

	var user models.User
	user.Name = "windness"
	user.Email = "windnessr@163.com"
	user.Password = util.Md5([]byte("123456"))
	user.Insert()

	beego.Run()
}
