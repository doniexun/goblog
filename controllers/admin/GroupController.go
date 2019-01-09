package admin

// GroupController 群组相关 controller
type GroupController struct {
	BaseController
}

// JoinGroup 加群
func (c *GroupController) JoinGroup() {
	//username := c.GetString("username")

}

// QuitGroup 退群
func (c *GroupController) QuitGroup() {
	//username := c.GetString("username")

}

// TransferGroup 转让群
func (c *GroupController) TransferGroup() {

}

// DissolveGroup 解散群
func (c *GroupController) DissolveGroup() {
	//user := models.User{ID: c.userid}

}

// UpdateAnnouncement 更新群公告
func (c *GroupController) UpdateAnnouncement() {
	// [TODO]
}

// UpdateAvatar 更新群头像
func (c *GroupController) UpdateAvatar() {
	// [TODO]
}

// AddMember 添加成员
func (c *GroupController) AddMember() {

}

// RemoveMember 移除成员
func (c *GroupController) RemoveMember() {

}
