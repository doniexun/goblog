package models

import (
	"github.com/astaxie/beego/orm"
)

// Option 系统配置项
type Option struct {
	ID    int64 `orm:"column(id);auto;pk"`
	Name  string
	Value string
}

// Insert 插入当前配置选项
func (m *Option) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

// Read 读取当前配置选项的多个字段
func (m *Option) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

// Update 更新当前配置选项的多个字段
func (m *Option) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

// Delete 删除当前配置选项
func (m *Option) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

// TableName 获取 TableName，需要调用 models/base.go 中定义的  TableName()
func (m *Option) TableName() string {
	return TableName("option")
}

// Query 获取 QuerySeter 对象
func (m *Option) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
