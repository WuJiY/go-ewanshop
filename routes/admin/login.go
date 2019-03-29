package admin

import (
	"github.com/haolifeng/go-ewanshop/framework"
)

func Login(ctx *framework.HandlerContext) {
	if ctx.R.Method == "GET" {
		var filename = "./views/admin/login.html"
		ctx.Render(filename, nil, nil)
		return
	}

	if ctx.R.Method == "POST" {
		var username = ctx.R.FormValue("username")
		var password = ctx.R.FormValue("password")
		var yzm = ctx.R.FormValue("yzm")

		var verifyOk = ctx.Verify(yzm)
		verifyOk = true
		if username == "admin" && password == "admin123" && verifyOk {
			ctx.SetSessionVal("isAdmin", true)
		}
		ctx.Redirect("./index.go")
	}
}
