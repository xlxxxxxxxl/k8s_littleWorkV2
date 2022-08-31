package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "ks8-demo/routers"
	//beego "github.com/astaxie/beego"
	//beego "github.com/astaxie/beego/context"
)

func main() {

	//_ = controllers.AddUser{}
	beego.Run()

	//beego.Get("/", func(ctx *context.Context) {
	//	// 用户数据的获取
	//	name := ctx.Input.Query("name")
	//
	//	// 给用户响应数据
	//	ctx.Output.Context.WriteString(fmt.Sprintf("您输入的名字是: %s", name))
	//})
	//
	//// Post
	//beego.Post("/", func(ctx *context.Context) {
	//	name := ctx.Input.Query("name")
	//	ctx.Output.Context.WriteString(fmt.Sprintf("(Post)您输入的名字是: %s", name))
	//})
	//
	//// 任意请求
	//beego.Any("/any", func(ctx *context.Context) {
	//	name := ctx.Input.Query("name")
	//	ctx.Output.Context.WriteString(fmt.Sprintf("(%s)您输入的名字是: %s", ctx.Input.Method(), name))
	//})
	//
	//// 启动Beego程序
	//beego.Run("127.0.0.1")
}
