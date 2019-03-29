package admin

import (
	"github.com/haolifeng/go-ewanshop/framework"
)

func Top(ctx *framework.HandlerContext) {
	if v, ok := ctx.GetSessionVal("isAdmin"); ok {
		if v.(bool) {
			ctx.HtmlFile("./views/admin/static/top.html")
			return
		}
	}
	ctx.Redirect("/routes/admin/login.go")
}
