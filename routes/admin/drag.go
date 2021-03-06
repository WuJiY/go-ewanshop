package admin

import (
	"github.com/haolifeng/go-ewanshop/framework"
)

func Drag(ctx *framework.HandlerContext) {
	if v, ok := ctx.GetSessionVal("isAdmin"); ok {
		if v.(bool) {
			ctx.HtmlFile("./views/admin/static/drag.html")
			return
		}
	}
	ctx.Redirect("/routes/admin/login.go")
}
