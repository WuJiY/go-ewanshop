package flow

import (
	"github.com/haolifeng/go-ewanshop/framework"
	"github.com/haolifeng/go-ewanshop/lib"
)

func Checkout(ctx *framework.HandlerContext) {
	var user = ctx.GetUser()
	if user == nil {
		ctx.Redirect("/routes/user/login.go")
		return
	}
	if ctx.R.Method == "GET" {
		var cart = ctx.GetCart()
		if len(cart.Items()) == 0 {
			ctx.Send("购物车为空")
			return
		}
		var allData = ctx.Cates.Find()
		var data = map[string]interface{}{
			"tree": ctx.Cates.GetTree(allData),
			"user": user,
			"cart": cart,
		}
		ctx.Render("./views/flow/checkout.html", data, lib.FuncMap)
		return
	}
}
