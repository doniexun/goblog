package admin

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
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
		group.ID = int64(groupID)
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
		group.NichName = "this 群组's nickname"
		group.Creator = user
		group.CreateIP = c.ClientIP()
		group.CreateTime = time.Now()
		group.LastUpdator = user
		group.LastUpdateIP = c.ClientIP()
		group.LastUpdateTime = time.Now()
		group.Announcement = "owner is lazy, no announcement..."
		if err := group.Insert(); err != nil {
			log.Println("群主信息插入DB失败" + err.Error())
			c.BackToClientReponse(false, "群主信息插入DB失败")
			return
		}

		// 群主与用户的多对多关系
		// 将用户添加到 group.Members 中
		// 在 group 表中插入对应 user 的 id
		// 在 user 表中并没有插入对应 group 的id
		m2m := orm.NewOrm().QueryM2M(group, "Members")
		if _, err := m2m.Add(user); err != nil {
			log.Println("用户信息插入群组失败" + err.Error())
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
		log.Println("打卡事项插入DB失败" + err.Error())
		c.BackToClientReponse(false, "打卡事项插入DB失败")
		return
	}

	// 打卡事项与打卡人的多对多关系
	m2m := orm.NewOrm().QueryM2M(user, "PunchItems")
	if _, err := m2m.Add(punchItem); err != nil {
		log.Println("打卡事项关联用户失败" + err.Error())
		c.BackToClientReponse(false, "打卡事项关联用户失败")
		return
	}

	// 打卡事项与群组的多对多关系
	// 将打卡事项添加到 group 中
	// 在 group 表中插入对应 punchItem 的 id
	// 在 punchItem 表中并没有插入对应 group 的id
	m2m = orm.NewOrm().QueryM2M(group, "PunchItems")
	if _, err := m2m.Add(punchItem); err != nil {
		log.Println("打卡事项插入群组失败" + err.Error())
		c.BackToClientReponse(false, "打卡事项插入群组失败")
		return
	}

	c.BackToClientReponse(true, "打卡事项插入群组成功")

}

// DeletePunch 删除打卡事项
func (c *PunchController) DeletePunch() {
	pid, err := c.GetInt64("id")
	if err != nil {
		log.Println("非法打卡事项" + err.Error())
		c.BackToClientReponse(false, "非法打卡事项")
		return
	}

	// 查询打卡事项
	var punchItem models.PunchItem
	punchItem.ID = pid
	if err = punchItem.Read(); err != nil {
		log.Println("打卡事项不存在" + err.Error())
		c.BackToClientReponse(false, "打卡事项不存在")
		return
	}

	title := punchItem.Title

	// 删除用户表中的记录
	// 删除多对多关系中的打卡事项
	user := models.User{ID: c.userID}
	m2m := orm.NewOrm().QueryM2M(&user, "PunchItems")
	if m2m.Exist(&punchItem) { // 只在存在 多对多关系时 删除；[TODO] 是否不判断，直接删除，效率更高？
		if _, err = m2m.Remove(&punchItem); err != nil {
			log.Println("用户表中的打卡事项移除失败" + err.Error())
			c.BackToClientReponse(false, "用户表中的打卡事项移除失败")
			return
		}
	}

	// 删除群组表中的记录
	// 删除多对多关系中的打卡事项
	if num, err := orm.NewOrm().LoadRelated(&punchItem, "Groups"); err == nil && num > 0 { // 加载关系字段
		for _, group := range punchItem.Groups {
			fmt.Println(group.Name)
			m2m := orm.NewOrm().QueryM2M(group, "PunchItems")
			if m2m.Exist(&punchItem) { // 只在存在 多对多关系时 删除；[TODO] 是否不判断，直接删除，效率更高？
				if _, err = m2m.Remove(&punchItem); err != nil {
					log.Println("群组 " + group.Name + " 中的打卡事项 " + title + "移除失败" + err.Error())
					c.BackToClientReponse(false, "群组 "+group.Name+" 中的打卡事项 "+title+"移除失败")
					return
				}
			}
		}
	}

	// 删除对应打卡记录表中的记录
	// 删除多对多关系中的打卡事项
	// 在 PunchRecord 表中设置了，当关联 PunchItem 删除时，自动级联删除
	// 所以此处无需代码操作

	// 删除打卡事项表中的记录
	if err = punchItem.Delete(); err != nil {
		log.Println("打卡事项 " + title + " 删除失败" + err.Error())
		c.BackToClientReponse(false, "打卡事项 "+title+" 删除失败")
		return
	}

	c.BackToClientReponse(true, "打卡事项 "+title+"删除成功")
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

// UserPunchList 获取当前登录用户的打卡事项清单
func (c *PunchController) UserPunchList() {

	var (
		err         error
		userID      int64
		punchItemID int64
	)

	punchItemID, err = c.GetInt64("punchid")
	if err != nil && strings.Contains(err.Error(), "parsing \"\": invalid syntax") {
		punchItemID = 0
	} else {
		log.Println("打卡事项编号有误，打卡事项清单查询失败" + err.Error())
		c.BackToClientReponse(false, "打卡事项编号有误，打卡事项清单查询失败")
		return
	}

	// 若用户未指定 punchid，则说明要查询的是打卡事项清单，否则要查询的是具体某个打卡事项信息
	if punchItemID == 0 { // 查询打卡事项清单
		// [WARN] 若用户输入的用户编号有误，也不应该为其查询当前登录用户的打卡事项清单
		//        因为这和用户本身意愿是不相符的
		userID, err = c.GetInt64("userid")
		if err != nil && strings.Contains(err.Error(), "parsing \"\": invalid syntax") {
			userID = c.userID // 用户未指定用户编号，则查询当前登录用户的打卡事项清单
		} else {
			log.Println("用户编号有误，打卡事项清单查询失败" + err.Error())
			c.BackToClientReponse(false, "用户编号有误，打卡事项清单查询失败")
			return
		}

		// 查询指定用户的存在
		user := &models.User{ID: userID}
		if err := user.Read(); err != nil {
			log.Println("指定用户不存在，打卡事项清单查询失败" + err.Error())
			c.BackToClientReponse(false, "指定用户不存在，打卡事项清单查询失败")
			return
		}

		var punchItems []*models.PunchItem
		_, err = orm.NewOrm().QueryTable("punch_item").Filter("Punchers__User__ID", userID).RelatedSel().All(&punchItems)
		if err == nil {
			c.BackToClientData(punchItems)
			return
		}

		c.BackToClientData("")
		return
	}

	// 查询具体打卡事项
	punchItem := &models.PunchItem{ID: punchItemID}
	err = punchItem.Query().Filter("ID", punchItemID).RelatedSel().One(punchItem)
	if err == nil {
		c.BackToClientData(punchItem)
		return
	}

	c.BackToClientData("")
	return
}
