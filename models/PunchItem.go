package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// PunchItem 打卡事项
type PunchItem struct {
	ID             int64        `orm:"column(id);auto;pk"`
	Title          string       `orm:"size(50)"`                             // 打卡事项标题
	Content        string       `orm:"size(500)"`                            // 打卡事项正文
	Creator        *User        `orm:"rel(fk)"`                              // 创建人；当某个用户被注销时，此项可能为空，因此用户注销时，一般只是设置其状态为“注销”状态，否则此处为 null 易出问题。
	CreateTime     time.Time    `orm:"auto_now_add;type(datetime)"`          // 创建时间
	CreateIP       string       `orm:"column(create_ip);size(50)"`           // 创建 IP 地址
	LastUpdator    *User        `orm:"rel(fk)"`                              // 最后更新人
	LastUpdateTime time.Time    `orm:"auto_now;type(datetime)"`              // 最后更新时间
	LastUpdateIP   string       `orm:"column(last_update_ip);size(50)"`      // 最后更新 IP 地址
	Status         int          `orm:"default(0)"`                           // 状态：0：活跃；1：关闭；2：删除
	Avatar         string       `orm:"size(100);null"`                       // 打卡封面头像所在 URL 地址
	BeginTime      time.Time    `orm:"type(datetime)"`                       // 开始打卡时间
	EndTime        time.Time    `orm:"type(datetime)"`                       // 结束打卡时间
	PeriodUnit     int          `orm:"size(5);default(3)"`                   // 周期的单位：0：秒，1：分钟；2：小时；3.天；4：周；5：月；6：季度；7：半年；8：年
	PeriodValue    int64        `orm:"default(1)"`                           // 打卡周期，默认一天周期
	ActiveBonus    int64        `orm:"default(0)"`                           // 活跃积分（活跃度）
	Groups         []*Group     `orm:"rel(m2m);null"`                        // 归属于群
	Punchers       []*User      `orm:"reverse(many);null"`                   // 关联的打卡人
	PunchRecord    *PunchRecord `orm:"reverse(one);null;on_delete(cascade)"` // 一个打卡事项对应一个打卡记录（在打卡事项删除时，级联删除）
}

// Insert 插入当前打卡事项
func (m *PunchItem) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

// Read 读取当前打卡事项的多个字段
func (m *PunchItem) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

// Update 更新当前打卡事项的多个字段
func (m *PunchItem) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

// Delete 删除当前打卡事项
func (m *PunchItem) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

// TableName 获取 TableName，需要调用 models/base.go 中定义的  TableName()
func (m *PunchItem) TableName() string {
	return TableName("punch_item")
}

// Query 获取 QuerySeter 对象
func (m *PunchItem) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
