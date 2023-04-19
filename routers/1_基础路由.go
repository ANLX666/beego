package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func init() {
	beego.Get("/get", func(context *context.Context) {
		//beego.Info("基础路由中的get请求")
		context.Output.Body([]byte("基础路由中的get请求 get method"))
	})
	beego.Post("/post", func(context *context.Context) {
		context.Output.Body([]byte("post"))
	})

}
