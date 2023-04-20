package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
)

// 定义User模型，绑定users表结构, 其实就是用来保存sql查询结果。
type User struct {
	Id       int
	Username string
	Password string
}

// 定义User 模型绑定那个表？
func (u *User) TableName() string {
	// 返回mysql表名
	return "users"
}

// 初始化函数，可以用来向orm注册model
func init() {
	// 向orm注册user模型
	orm.RegisterModel(&User{})
}

// 根据id查询用户信息
func GetUserById(id int) *User {
	if id == 0 {
		return nil
	}

	// 创建orm对象, 后面都是通过orm对象操作数据库
	o := orm.NewOrm()

	// 初始化一个User模型对象
	user := User{}
	// 设置查询参数
	user.Id = id

	// 调用Read方法，根据user设置的参数，查询一条记录，结果保存到user结构体变量中
	// 默认是根据主键进行查询
	// 等价sql： SELECT `id`, `username`, `password` FROM `users` WHERE `id` = 1
	err := o.Read(&user)

	// 检测查询结果，
	if err == orm.ErrNoRows {
		// 找不到记录
		return nil
	} else if err == orm.ErrMissPK {
		// 找不到住建
		return nil
	}

	return &user
}
func GetCount() int64 {
	fmt.Println(1)
	o := orm.NewOrm()
	cnt, err := o.QueryTable("users").Count() // SELECT COUNT(*) FROM USER
	fmt.Printf("Count Num: %s, %s", cnt, err)
	return cnt
}
func GetUserArr() []User {
	o := orm.NewOrm()
	arr := o.QueryTable("users")
	var users []User
	all, _ := arr.All(&users)
	fmt.Println(all)
	return users
}
func DeleteUser(id string) int64 {
	o := orm.NewOrm()
	i, _ := o.QueryTable("users").Filter("id", id).Delete()
	return i
}
