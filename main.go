package main

import (
	"github.com/astaxie/beego"
	"github.com/doniexun/goblog/models"
	_ "github.com/doniexun/goblog/routers"
)

func main() {

	var post models.Post
	post.AuthorId = 1
	post.AuthorName = "windness"
	post.Title = "Test Post Title"
	post.Content = "A testing post's content"

	post.Insert()

	beego.Run()
}
