package models

// 如果用户利用用户菜单创建打卡（事项）时，
// 会创建一个 PunchItem ，
// 同时会创建一个 Group ，并自动成为群主、群成员之一、创建者，
// 并将该 PunchItem 绑定到 Group 中。

// 如果用户利用群菜单创建打卡（事项）时，
// 会创建一个 PunchItem ,
// 并将该 PunchItem 绑定到 Group 中。

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // 用于引入 mysql 驱动
)

func init() {

	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname := beego.AppConfig.String("dbname")

	// 设置默认时区，否则无法在数据表中存储中文
	// 切必须在注册数据库语句 orm.RegisterDatabase() 调用之前
	orm.DefaultTimeLoc = time.Local

	//注册mysql Driver
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//构造conn连接
	//用户名:密码@tcp(url地址)/数据库
	conn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8&loc=Asia%2FShanghai"
	//注册default数据库
	orm.RegisterDataBase("default", "mysql", conn)
	fmt.Printf("数据库连接成功！%s\r\n", conn)

	// 注册数据库表（若不注册模型，则不会在DB中生成对应的表）
	orm.RegisterModel(
		new(User),
		new(Group),
		new(Post),
		new(Option),
		new(PunchItem),
		new(PunchRecord))

	//orm.RegisterModelWithPrefix("t_", new(User), new(Post), new(Option))  //带前缀的表

	// 只在开发模式下才开启 orm 的 Debug 功能
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}

	orm.RunSyncdb("default", false, true) // true 改成false，如果表存在则会给出提示，如果改成false则不会提示 ，这句话没有会报主键不存在的错误

}

// TableName 返回数据表名
func TableName(str string) string {
	return beego.AppConfig.String("dbprefix") + str
}
