package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Option struct {
	Id		int64		`orm:"auto;pk"`
	Name		string
	Value		string
}

/**
 * 插入当前配置选项
 */
func (m *Option) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

/**
 * 读取当前配置选项的多个字段
 */
func (m *Option) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err 
	}
	return nil
}

/**
 * 更新当前配置选项的多个字段
 */
func (m *Option) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

/**
 * 删除当前配置选项
 */
func (m *Option) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

/**
 * 获取 TableName，需要调用 models/base.go 中定义的  TableName()
 */
func (m *Option) TableName() string {
	return TableName("option")
}


/**
 * 获取 QuerySeter 对象
 */
func (m *Option) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}


