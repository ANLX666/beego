package main

import (
	_ "beegoProject/routers"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	//导入mysql驱动，这是必须的
	_ "github.com/go-sql-driver/mysql"
)

// 初始化应用设置， 我们通过init函数初始化数据库连接，go语言中这个函数会优先执行
func init() {
	// 这里注册一个default默认数据库，数据库驱动是mysql.
	// 第三个参数是数据库dsn, 配置数据库的账号密码，数据库名等参数
	//  dsn参数说明：
	//      username    - mysql账号
	//      password    - mysql密码
	//      db_name     - 数据库名
	//      127.0.0.1:3306 - 数据库的地址和端口
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/beegoblog?charset=utf8")
}

func main() {
	beego.Run()
}
