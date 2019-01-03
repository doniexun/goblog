package models

import (
	"github.com/astaxie/beego/orm"
)

// User 用户
type User struct {
	ID       int64  `orm:"auto;pk"`
	Name     string `orm:"unique;size(100)"`
	Password string `orm:"size(100)"`
	Email    string `orm:"unique;size(100)"`
}

// Insert 插入当前用户
func (m *User) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

// Read 读取当前用户的多个字段
func (m *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

// Update 更新当前用户的多个字段
func (m *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

// Delete 删除当前用户
func (m *User) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

// TableName 获取 TableName，需要调用 models/base.go 中定义的  TableName()
func (m *User) TableName() string {
	return TableName("user")
}

// Query 获取 QuerySeter 对象
func (m *User) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
