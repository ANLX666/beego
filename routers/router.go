package routers

import (
	"beegoProject/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {

	beego.Router("/user/:id", &controllers.MainController{})
	beego.Router("/*", &controllers.CountController{})
	beego.Router("/index", &controllers.IndexController{})
	beego.Router("/getUserList", &controllers.GetUserList{})
	beego.Router("/deleteUser/:id", &controllers.DeleteUser{})
}
