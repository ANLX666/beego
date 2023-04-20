package controllers

import (
	"beegoProject/models"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"strconv"
)

type MainController struct {
	beego.Controller
}
type CountController struct {
	beego.Controller
}
type IndexController struct {
	beego.Controller
}
type GetUserList struct {
	beego.Controller
}
type DeleteUser struct {
	beego.Controller
}
type AddUser struct {
	beego.Controller
}
type RegisterMain struct {
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

func (c *CountController) Get() {
	c.Data["Website"] = "1209322734.com"
	c.Data["Email"] = "1209322734@demo.com"
	fmt.Println(1)
	count := models.GetCount()
	c.Controller.Data["json"] = count
	c.ServeJSON()
	c.TplName = "user.tpl"
}
func (c *IndexController) Post() {
	c.Data["test"] = "区块链"
	c.TplName = "test.html"
}
func (c *GetUserList) Get() {
	UserList := models.GetUserArr()
	c.Controller.Data["json"] = UserList
	c.ServeJSON()
	c.TplName = "user.tpl"
}
func (c *DeleteUser) Delete() {
	id := c.Ctx.Input.Param(":id")
	user := models.DeleteUser(id)
	c.Controller.Data["json"] = user
	c.ServeJSON()
	c.TplName = "user.tpl"
}
func (c *AddUser) Post() {
	username := c.Ctx.Input.Param(":username")
	password := c.Ctx.Input.Param(":password")

	user := models.AddUser(username, password)
	c.Controller.Data["json"] = user
	c.ServeJSON()
	c.TplName = "user.tpl"
}
func (c *RegisterMain) PostRegister() {
	userName := c.GetStrings("userName")
	password := c.GetStrings("password")
	o := orm.NewOrm()
	user := models.User{}
	fmt.Println(userName)
	fmt.Println(password)

	user.Username = userName[0]
	user.Password = password[0]
	_, err := o.Insert(&user)
	if err != nil {
		fmt.Println(err)
		c.Redirect("/register", 302)
	}
	c.Ctx.WriteString("注册成功")
}
func (c *RegisterMain) GetRegister() {

	c.TplName = "register.html"
}
