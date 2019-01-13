package admin

import (
	"fmt"

	"github.com/doniexun/goblog/models"
)

// PostController 用户账户 controller
type PostController struct {
	BaseController
}

// List 指定页的文章列表
func (c *PostController) List() {
	var (
		list     []*models.Post
		err      error
		page     int
		pagesize int
	)

	if page, err = c.GetInt("page"); err != nil || page < 1 {
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
	c.TplName = "admin/index.tpl" // 相对路径：相对于工程根目录 views/
}

// Add 添加文章
func (c *PostController) Add() {
	var (
		post models.Post
	)

	title := c.GetString("title")
	content := c.GetString("content")

	if title == "" {
		c.Data["errmsg"] = "标题不能为空"
		fmt.Println("标题不能为空")
		return
	}

	post.AuthorID = c.userID
	post.AuthorName = c.userName
	post.Content = content
	post.Title = title
	post.Insert()

	// 添加成功后向客户端返回结果
	// ...[TODO]
}

// Delete 删除文章
func (c *PostController) Delete() {
	id, _ := c.GetInt64("id")
	post := models.Post{ID: id}
	if post.Read() == nil {
		post.Delete()
	}
	c.Redirect("/admin/post/list", 302)
}

// Update 修改文章
func (c *PostController) Update() {
	var (
		id   int
		post models.Post
	)

	id, _ = c.GetInt("id")
	post.ID = int64(id)
	if post.Read() != nil {
		c.Abort("404")
	}

	title := c.GetString("title")
	content := c.GetString("content")

	if title == "" {
		c.Data["errmsg"] = "标题不能为空"
		fmt.Println("标题不能为空")
		return
	}

	post.AuthorID = c.userID
	post.AuthorName = c.userName
	post.Content = content
	post.Title = title
	post.Update("title", "content")

	// 添加成功后向客户端返回结果
	// ...[TODO]
	c.Redirect("/admin/post/list", 302)
}

// AddDemoPost 添加示例文章
func (c *PostController) AddDemoPost() {
	var post models.Post
	post.AuthorID = 1
	post.AuthorName = "windness"
	post.Title = "Test Post Title"
	post.Content = "A testing post's content"
	post.Insert()
}
