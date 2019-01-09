package admin

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/doniexun/goblog/models"
	"github.com/lisijie/goblog/util"
)

// PunchController 打卡相关 controller
type PunchController struct {
	BaseController
}

// CreatePunch 创建打卡事项
func (c *PunchController) CreatePunch() {

	punchItem := &models.PunchItem{}
	punchItem.Title = "test title"
	punchItem.Content = "test content"
	punchItem.CreateTime = time.Now()
	punchItem.CreateIP = c.ClientIP()
	punchItem.LastUpdateTime = time.Now()
	punchItem.LastUpdateIP = c.ClientIP()
	punchItem.BeginTime = time.Now()
	duration, _ := time.ParseDuration("+24h")
	punchItem.EndTime = time.Now().Add(duration)
	punchItem.PeriodUnit = 3
	punchItem.PeriodValue = 1

	user := &models.User{}
	user.UserName = "windness"
	user.NickName = "doniexun"
	user.Password = util.Md5([]byte("123456"))
	user.Email = "windnessr@163.com"
	user.Cellphone = "13800138000"
	user.QQ = "1758953369"
	user.Wechat = "splendidream"
	user.RegisterTime = time.Now()
	user.RegsiterIP = c.ClientIP()
	user.LastLoginTime = time.Now()
	user.LastLoginIP = c.ClientIP()
	if err := user.Insert(); err != nil {
		fmt.Println(err.Error())
	}

	punchItem.Creator = user
	punchItem.LastUpdator = user
	if err := punchItem.Insert(); err != nil {
		fmt.Println(err.Error())
	}

	// 在 punch_item 表中插入对应 user 的 id
	// 在 user 表中并没有插入对应 punch_item 的id
	m2m := orm.NewOrm().QueryM2M(punchItem, "punchers")
	if num, err := m2m.Add(user); err != nil {
		fmt.Println("num ----- ", num)
	} else {
		fmt.Println(err.Error())
	}

	// if c.userid < 0 {
	// 	return
	// }

	// title := c.GetString("title")
	// content := c.GetString("content")
	// begintimestr := c.GetString("begintime")
	// endtimestr := c.GetString("endtime")
	// periodstr := c.GetString("period") // 打卡周期，默认24小时

	// begintime, _ := time.Parse("2006-01-02 15:04:05", begintimestr)
	// endtime, _ := time.Parse("2006-01-02 15:04:05", endtimestr)
	// period, _ := time.Parse("2006-01-02 15:04:05", periodstr)

	// var punchItem models.PunchItem
	// punchItem.Title = title
	// punchItem.Content = content
	// punchItem.BeginTime = begintime
	// punchItem.EndTime = endtime
	// punchItem.Period = period

	// var user models.User
	// user.ID = c.userid
	// if err := user.Read(); err != nil {
	// 	return
	// }

	// // 根据传回的 type 类型来判定是来自个人菜单中的打卡创建，还是来自群的打卡创建
	// typeValue, _ := c.GetInt("type")
	// if typeValue == 0 { // 若是个人菜单创建
	// 	// 新建一个 Group

	// } else { // 若是群菜单创建
	// 	// 查找 Group
	// 	groupID, _ := c.GetInt("groupid")
	// 	var group models.Group
	// 	group.ID = groupID
	// 	if err := group.Read(); err != nil {
	// 		return
	// 	}

	// 	// 判断当前用户是否已经加入群
	// 	// 若未加群，则应该报错：非本群用户不允许加入此群
	// 	// [TODO]
	// 	// 正常来说是不应该存在这个问题的

	// 	// 判断当前群是否已经存在同名打卡事项
	// 	// 同一个群中不允许存在同名打卡事项
	// 	// [TODO]

	// 	// 将打卡事项、群、用户相互关联起来
	// 	// 打卡事项与群

	// 	// 打卡事项与用户

	// }

}

// DeletePunch 删除打卡事项
func (c *PunchController) DeletePunch() {
	//username := c.GetString("username")

}

// ClosePunch 关闭打卡事项
func (c *PunchController) ClosePunch() {

}

// ResumePunch 恢复打卡事项
func (c *PunchController) ResumePunch() {
	//user := models.User{ID: c.userid}

}

// ClearPunchRecord 清空打卡记录
func (c *PunchController) ClearPunchRecord() {
	// [TODO]
}

// UpdatePunchAvatar 更新打卡事项头像
func (c *PunchController) UpdatePunchAvatar() {
	// [TODO]
}

// JoinPunch 加入打卡事项
func (c *PunchController) JoinPunch() {
	// [TODO]
}

// QuitPunch 退出打卡事项
func (c *PunchController) QuitPunch() {
	// [TODO]
}

// PunchInfo 获取打卡事项信息
func (c *PunchController) PunchInfo() {

}

// PunchList 获取打卡事项清单
func (c *PunchController) PunchList() {

}
