package main

import (
	_ "github.com/doniexun/goblog/routers"
	"github.com/doniexun/goblog/models"
	"github.com/astaxie/beego"
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

