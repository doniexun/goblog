package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Id		int64		`orm:"auto;pk"`
	AuthorId	int64
	AuthorName	string		`orm:"size(128)"`
	Title		string		`orm:"size(256)"`
	Content		string		`orm:"type(text)"`
}

/**
 * 插入当前文章
 */
func (m *Post) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

/**
 * 读取当前文章的多个字段
 */
func (m *Post) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err 
	}
	return nil
}

/**
 * 更新当前文章的多个字段
 */
func (m *Post) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

/**
 * 删除当前文章
 */
func (m *Post) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

/**
 * 获取 TableName，需要调用 models/base.go 中定义的  TableName()
 */
func (m *Post) TableName() string {
	return TableName("post")
}


/**
 * 获取 QuerySeter 对象
 */
func (m *Post) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}


