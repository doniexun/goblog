package admin

// PunchController 打卡相关 controller
type PunchController struct {
	BaseController
}

// CreatePunch 创建打卡事项
func (c *PunchController) CreatePunch() {
	//username := c.GetString("username")

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
