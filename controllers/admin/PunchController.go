package admin

import (
	"time"

	"github.com/doniexun/goblog/models"
)

// PunchController 打卡相关 controller
type PunchController struct {
	BaseController
}

// CreatePunch 创建打卡事项
func (c *PunchController) CreatePunch() {

	if c.userid < 0 {
		return
	}

	title := c.GetString("title")
	content := c.GetString("content")
	begintimestr := c.GetString("begintime")
	endtimestr := c.GetString("endtime")
	periodstr := c.GetString("period") // 打卡周期，默认24小时

	begintime, _ := time.Parse("2006-01-02 15:04:05", begintimestr)
	endtime, _ := time.Parse("2006-01-02 15:04:05", endtimestr)
	period, _ := time.Parse("2006-01-02 15:04:05", periodstr)

	var punchItem models.PunchItem
	punchItem.Title = title
	punchItem.Content = content
	punchItem.BeginTime = begintime
	punchItem.EndTime = endtime
	punchItem.Period = period

	var user models.User
	user.ID = c.userid
	if err := user.Read(); err != nil {
		return
	}

	// 根据传回的 type 类型来判定是来自个人菜单中的打卡创建，还是来自群的打卡创建
	typeValue, _ := c.GetInt("type")
	if typeValue == 0 { // 若是个人菜单创建
		// 新建一个 Group

	} else { // 若是群菜单创建
		// 查找 Group
		groupID, _ := c.GetInt("groupid")
		var group models.Group
		group.ID = groupID
		if err := group.Read(); err != nil {
			return
		}

		// 判断当前用户是否已经加入群
		// 若未加群，则应该报错：非本群用户不允许加入此群
		// [TODO]
		// 正常来说是不应该存在这个问题的

		// 判断当前群是否已经存在同名打卡事项
		// 同一个群中不允许存在同名打卡事项
		// [TODO]

		// 将打卡事项、群、用户相互关联起来
		// 打卡事项与群

		// 打卡事项与用户

	}

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
