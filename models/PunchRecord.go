package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// PunchRecord 打卡记录
type PunchRecord struct {
	ID        int        `orm:"column(id);auto;pk"`
	Puncher   *User      `orm:"rel(fk)"`                     // 打卡人
	PunchItem *PunchItem `orm:"rel(one)"`                    // 一个打卡记录只能对应一个打卡事项，一个打卡事项也只能对应一个打卡记录
	PunchTime time.Time  `orm:"auto_now_add;type(datetime)"` // 打卡时间

}

// Insert 插入当前打卡记录
func (m *PunchRecord) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

// Read 读取当前打卡记录的多个字段
func (m *PunchRecord) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

// Update 更新当前记录的多个字段
func (m *PunchRecord) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

// Delete 删除当前记录事项
func (m *PunchRecord) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

// TableName 获取 TableName，需要调用 models/base.go 中定义的  TableName()
func (m *PunchRecord) TableName() string {
	return TableName("punch_record")
}

// Query 获取 QuerySeter 对象
func (m *PunchRecord) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
