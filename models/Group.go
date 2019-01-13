package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// Group 群
type Group struct {
	ID             int64        `orm:"column(id);auto;pk"`
	Name           string       `orm:"size(50)"`                        // 群名称
	NichName       string       `orm:"size(50)"`                        // 备注群名称
	Owner          *User        `orm:"rel(fk)"`                         // 群主（一个群只能有一个群主，群主与群的关系是一对一的关系）
	Creator        *User        `orm:"rel(fk)"`                         // 创建人（一个群只能有一个创建人，创建人与群的关系是一对一的关系）
	CreateTime     time.Time    `orm:"auto_now_add;type(datetime)"`     // 创建时间
	CreateIP       string       `orm:"column(create_ip);size(50)"`      // 创建 IP 地址
	LastUpdator    *User        `orm:"rel(fk)"`                         // 最后更新人
	LastUpdateTime time.Time    `orm:"auto_now;type(datetime)"`         // 最后更新时间
	LastUpdateIP   string       `orm:"column(last_update_ip);size(50)"` // 最后更新 IP 地址
	ActiveBonus    int          `orm:"default(0)"`                      // 活跃积分
	Status         int          `orm:"default(0)"`                      // 群账号状态：0：正常；1：异常；2：注销
	Announcement   string       `orm:"size(512);null"`                  // 群公告
	Members        []*User      `orm:"reverse(many);null"`              // 群成员（与群是多对多的关系）
	PunchItems     []*PunchItem `orm:"reverse(many);null"`              // 打卡事项（一个群可以有多个打卡事项，对应多个打卡记录，群和打卡事项是多对多的关系）
}

// Insert 插入当前群
func (m *Group) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

// Read 读取当前群的多个字段
func (m *Group) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

// Update 更新当前群的多个字段
func (m *Group) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

// Delete 删除当前群
func (m *Group) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

// TableName 获取 TableName，需要调用 models/base.go 中定义的  TableName()
func (m *Group) TableName() string {
	return TableName("group")
}

// Query 获取 QuerySeter 对象
func (m *Group) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
