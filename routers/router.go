package routers

import (
	"github.com/astaxie/beego"
	"github.com/doniexun/goblog/controllers/admin"
	"github.com/doniexun/goblog/controllers/front"
)

func init() {
	// 访问示例：GET /
	beego.Router("/", &front.MainController{}, "get:Index") // 前台首页（显示指定页的文章）
	// 访问示例：GET /post/1
	beego.Router("/post/:id:int", &front.MainController{}, "get:Show") // 显示指定文章

	// [调试]添加demo用户接口：POST /test/adddemouser
	beego.Router("/test/adddemouser", &admin.AccountController{}, "post:AddDemoUser")
	// [调试]添加demo文章接口：POST /test/adddemopost
	beego.Router("/test/adddemopost", &admin.PostController{}, "post:AddDemoPost")
	// [调试]创建打卡事项接口：POST /test/createpunch[?group=xxx]
	beego.Router("/test/createpunch", &admin.PunchController{}, "post:CreatePunch")
	// [调试]删除打卡事项接口：GET /test/deletepunch?id=xxx
	beego.Router("/test/deletepunch", &admin.PunchController{}, "get:DeletePunch")
	// [调试]查询用户所参与的打卡事项列表[指定编号打卡事项信息]，数据以 json 格式返回：GET /test/userpunchs[?userid=xxx&punchid=xxx]
	//      若提供了 userid、punchid，则返回指定用户的指定打卡事项的详细信息
	//      若未提供 userid 、 punchid，则返回当前登录用户的打卡事项列表
	//      若提供了 userid，未提供 punchid，则返回指定用户的打卡事项列表
	//      若未提供 userid，提供了 punchid，则返回当前用户指定打卡事项的详细信息
	// [TODO] 未带其他过滤参数，如活跃状态、指定时间段、参与人数最多……等过滤参数
	// [TODO] 增加对查询打卡事项的权限，只有在权限表中的用户才有查询权限
	beego.Router("/test/userpunchs", &admin.PunchController{}, "get:Punchs")
	// [调试]获取服务器返回的json格式数据：GET /test/json
	beego.Router("/test/json", &admin.PunchController{}, "get:JSONInfo")

	// 访问示例：POST /admin/account/register，参数：username=xxx,password=xxx
	beego.Router("/admin/account/register", &admin.AccountController{}, "post:Register") // 注册
	// 访问示例：POST /admin/account/login，参数：username=xxx,password=xxx
	beego.Router("/admin/account/login", &admin.AccountController{}, "post:Login") // 登录
	// 访问示例：GET /admin/account/logout
	beego.Router("/admin/account/logout", &admin.AccountController{}, "get:Logout") // 登出
	// 访问示例：GET /admin/account/profile?id=1
	beego.Router("/admin/account/profile", &admin.AccountController{}, "get:Profile") // 个人信息

	// 访问示例：GET /admin/post/list?page=2
	// 访问示例：GET /admin/post/list
	beego.Router("/admin/post/list", &admin.PostController{}, "get:List") // 显示指定页的文章
	// 访问示例：POST /admin/post/add，参数：title=xxx,content=xxx etc.
	beego.Router("/admin/post/add", &admin.PostController{}, "post:Add") // 添加文章
	// 访问示例：GET /admin/post/delete?id=xx
	beego.Router("/admin/post/delete", &admin.PostController{}, "get:Delete") // 删除文章
	// 访问示例：POST /admin/post/update，参数：id=xxx,title=xxx,content=xxx
	beego.Router("/admin/post/update", &admin.PostController{}, "post:Update") // 更新文章
}
