package user

import (
	"github.com/haolifeng/go-ewanshop/framework"
	"github.com/haolifeng/go-ewanshop/lib"
)

func Info(ctx *framework.HandlerContext) {
	var user = ctx.GetUser()
	if user == nil {
		ctx.Redirect("/routes/user/login.go")
		return
	}
	if ctx.R.Method == "GET" {
		var allData = ctx.Cates.Find()
		var data = map[string]interface{}{
			"tree": ctx.Cates.GetTree(allData),
			"user": user,
			"cart": ctx.GetCart(),
		}
		ctx.Render("./views/user/info.html", data, lib.FuncMap)
		return
	}
}
