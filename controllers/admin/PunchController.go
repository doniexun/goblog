package admin

import (
	"encoding/json"
	"log"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/doniexun/goblog/models"
)

// PunchController 打卡相关 controller
type PunchController struct {
	BaseController
}

// JSONInfo 返回 Json 格式消息测试
func (c *PunchController) JSONInfo() {
	//content := `{"Name":"windness", "Age":20}`
	//c.Ctx.WriteString(content)
	//c.BackToClientReponse(true, "Hello World!\n")

	user := &models.User{}
	user.ID = 1
	if err := user.Read(); err == nil {
		c.BackToClientData(user)
	}
}

// CreatePunch 创建打卡事项
/// 从客户端 POST 来的 json 数据只包括：
//// Title    打卡标题
//// Content  打卡详细描述
//// BeginTime  开始打卡时间
//// EndTime    结束打卡时间
//// PeriodUnit   打卡周期单位
//// PeriodValue  打卡周期数值
/// 从接口中获取的参数包括：
//// group 群组编号
/// [TODO] 待修改成 orm 事件处理
func (c *PunchController) CreatePunch() {

	// 解析 json 数据，并创建 打卡事项 PunchItem 实例
	var punchItem models.PunchItem
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &punchItem); err != nil {
		c.BackToClientReponse(false, "提交信息无法识别，请重新提交！"+err.Error())
		return
	}
	punchItem.CreateTime = time.Now()
	punchItem.CreateIP = c.ClientIP()
	punchItem.LastUpdateTime = time.Now()
	punchItem.LastUpdateIP = c.ClientIP()

	// 获取当前登录用户信息
	user := &models.User{}
	user.ID = c.userID
	if err := user.Read(); err != nil {
		c.BackToClientReponse(false, "当前用户未登录"+err.Error())
		return
	}

	groupID, err := c.GetInt("group")
	if err != nil {
		groupID = 0 // 若从请求中未获取到群编号
	}

	var group *models.Group
	if groupID > 0 { // 群内创建
		// 查找群实例
		group = new(models.Group)
		group.ID = groupID
		if err := group.Read(); err != nil {
			log.Println("群实例未能找到，原因可能是群编号有误" + err.Error())
			c.BackToClientReponse(false, "群实例未能找到，原因可能是群编号有误")
			return
		}
	} else { // 个人创建
		// 创建群实例
		group = new(models.Group)
		group.Owner = user
		group.Name = user.UserName + "'s group"
		group.NichName = "this group's nickname"
		group.Creator = user
		group.CreateIP = c.ClientIP()
		group.CreateTime = time.Now()
		group.LastUpdator = user
		group.LastUpdateIP = c.ClientIP()
		group.LastUpdateTime = time.Now()
		group.Announcement = "owner is lazy, no announcement..."
		if err := group.Insert(); err != nil {
			log.Println(err.Error())
			c.BackToClientReponse(false, "群主信息插入DB失败")
			return
		}

		// 将用户添加到 group.Members 中
		// 在 group 表中插入对应 user 的 id
		// 在 user 表中并没有插入对应 group 的id
		m2m := orm.NewOrm().QueryM2M(group, "Members")
		if _, err := m2m.Add(user); err != nil {
			log.Println(err.Error())
			c.BackToClientReponse(false, "用户信息插入群组失败")
			return
		}
	}

	// 关联关系
	// 打卡事项与创建者、更新者关系
	// 注意： User 中不记录打卡事项，只有在 User 打卡时，User中才会有 PunchRecords 的记录，但不会有 PunchItems 的关联记录
	punchItem.Creator = user
	punchItem.LastUpdator = user
	if err := punchItem.Insert(); err != nil {
		log.Println(err.Error())
		c.BackToClientReponse(false, "打卡事项插入DB失败")
		return
	}

	// 打卡事项与群组的关系
	// 将打卡事项添加到 group 中
	// 在 group 表中插入对应 punchItem 的 id
	// 在 punchItem 表中并没有插入对应 group 的id
	m2m := orm.NewOrm().QueryM2M(group, "PunchItems")
	if _, err := m2m.Add(punchItem); err != nil {
		log.Println(err.Error())
		c.BackToClientReponse(false, "打卡事项插入群组失败")
		return
	}

	c.BackToClientReponse(true, "打卡事项插入群组成功")

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
