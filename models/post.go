package models

import (
	"github.com/astaxie/beego/orm"
)

// Post 文章
type Post struct {
	ID         int64 `orm:"auto;pk"`
	AuthorID   int64
	AuthorName string `orm:"size(128)"`
	Title      string `orm:"size(256)"`
	Content    string `orm:"type(text)"`
}

// Insert 插入当前文章
func (m *Post) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

// Read 读取当前文章的多个字段
func (m *Post) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

// Update 更新当前文章的多个字段
func (m *Post) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

// Delete 删除当前文章
func (m *Post) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

// TableName 获取 TableName，需要调用 models/base.go 中定义的  TableName()
func (m *Post) TableName() string {
	return TableName("post")
}

// Query 获取 QuerySeter 对象
func (m *Post) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
