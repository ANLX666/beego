package controllers

import (
	"beegoProject/models"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"strconv"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "1209322734.com"
	c.Data["Email"] = "1209322734@demo.com"

	id := c.Ctx.Input.Param(":id")

	// 调用model，查询用户id为1 的用户	信息

	zhuanhuahoudeId, _ := strconv.Atoi(id)
	user := models.GetUserById(zhuanhuahoudeId)
	// 然后将user数据保存到Data中, 将参数传给后面的views视图模板处理

	c.Controller.Data["json"] = user
	c.ServeJSON()
	fmt.Println(user)
	//c.Data["user"] = user
	// 使用新的视图模板user.tpl
	c.TplName = "user.tpl"
}
