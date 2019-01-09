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

	// 调试用接口：POST /test/adddemouser
	beego.Router("/test/adddemouser", &admin.AccountController{}, "post:AddDemoUser")
	// 调试用接口：POST /test/adddemopost
	beego.Router("/test/adddemopost", &admin.PostController{}, "post:AddDemoPost")

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
