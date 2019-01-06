package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// Punch 打卡事项
type Punch struct {
	ID             int       `orm:"column(id);auto;pk"`
	Title          string    `orm:"size(50)"`                    // 打卡事项标题
	Content        string    `orm:"size(500)"`                   // 打卡事项正文
	Creator        *User     `orm:"rel(fk)"`                     // 创建人
	CreateTime     time.Time `orm:"auto_now_add;type(datetime)"` // 创建时间
	CreateIP       string    `orm:"size(50)"`                    // 创建 IP 地址
	LastUpdator    *User     `orm:"rel(fk)"`                     // 最后更新人
	LastUpdateTime time.Time `orm:"auto_now_add;type(datetime)"` // 最后更新时间
	LastUpdateIP   string    `orm:"size(50)"`                    // 最后更新 IP 地址
	Status         int       `orm:"default(0)"`                  // 状态：0：活跃；1：关闭；2：删除
	Avatar         string    `orm:"size(100);null"`              // 打卡封面头像所在 URL 地址
	BeginTime      time.Time `orm:"type(datetime)"`              // 开始打卡时间
	EndTime        time.Time `orm:"type(datetime)"`              // 结束打卡时间
	Period         time.Time `orm:"type(datetime);null"`         // 打卡周期 [TODO] 默认一天周期
	Groups         []*Group  `orm:"rel(many)"`                   // 所属群（一个事项可以存在与多个群中，打卡事项与群的关系是多对多的关系）
	ActiveBonus    int       `orm:"default(0)"`                  // 活跃积分（活跃度）
}

// Insert 插入当前打卡事项
func (m *Punch) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

// Read 读取当前打卡事项的多个字段
func (m *Punch) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

// Update 更新当前用户的多个字段
func (m *Punch) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

// Delete 删除当前打卡事项
func (m *Punch) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

// TableName 获取 TableName，需要调用 models/base.go 中定义的  TableName()
func (m *Punch) TableName() string {
	return TableName("user")
}

// Query 获取 QuerySeter 对象
func (m *Punch) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
