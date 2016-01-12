package main

import (
	_ "archipel_admin_api/docs"
	_ "archipel_admin_api/routers"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego"
)

func main() {
	beego.InsertFilter("docs/*", beego.BeforeRouter, func(ctx *context.Context) {
    		if ctx.Input.Method() == "OPTIONS" {
			ctx.Output.Header("Access-Control-Allow-Origin", "*")
			ctx.Output.Header("Access-Control-Allow-Methods", "GET,POST,DELETE,PUT")
			ctx.Output.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
			ctx.Abort(200, "Hello")
		}
	})
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
