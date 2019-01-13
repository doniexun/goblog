package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// User 用户
/// [TODO] User 还是应该与 PunchItem 直接关联，不然在还未打卡前，无法查询用户加入的打卡事项。
type User struct {
	ID            int64          `orm:"column(id);auto;pk"`
	UserName      string         `orm:"unique;size(100)"`               // 用户名
	NickName      string         `orm:"unique;size(100);null"`          // 用户昵称
	Password      string         `orm:"size(100)"`                      // 用户密码
	Avatar        string         `orm:"size(100);null"`                 // 用户头像所在 URL 地址
	Email         string         `orm:"unique;size(100)"`               // 用户邮箱
	QQ            string         `orm:"column(qq);size(20);null"`       // QQ号
	Wechat        string         `orm:"size(255);null"`                 // 微信号
	Cellphone     string         `orm:"size(20);null"`                  // 手机号
	RegisterTime  time.Time      `orm:"auto_now_add;type(datetime)"`    // 注册时间
	RegsiterIP    string         `orm:"column(register_ip);size(50)"`   // 注册 IP 地址
	LastLoginTime time.Time      `orm:"type(datetime)"`                 // 最后登录时间
	LastLoginIP   string         `orm:"column(last_login_ip);size(50)"` // 最后登录 IP 地址
	LevelBonus    int            `orm:"default(0)"`                     // 等级积分
	ActiveBonus   int            `orm:"default(0)"`                     // 活跃积分
	Status        int            `orm:"default(0)"`                     // 用户账号状态：0：正常；1：异常；2：注销
	Groups        []*Group       `orm:"rel(m2m)"`                       // 用户所在群（若要查该用户创建的群，需要二次查询群表）
	PunchRecords  []*PunchRecord `orm:"reverse(many)"`                  // 用户的打卡记录
	PunchItems    []*PunchItem   `orm:"rel(m2m)"`                       // 用户关注的打卡事项
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
