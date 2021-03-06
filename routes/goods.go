package routes

import (
	"github.com/haolifeng/go-ewanshop/framework"
	"github.com/haolifeng/go-ewanshop/lib"
)

func Goods(ctx *framework.HandlerContext) {

	if ctx.R.Method == "GET" {
		var oid = ctx.R.FormValue("oid")
		var gs = ctx.Goods.GetByOid(oid)
		if gs == nil {
			ctx.Send("goods not found")
			return
		}
		ctx.AddHistroy(gs)
		var cats = ctx.Cates.Find()
		var data = map[string]interface{}{
			"gs":     gs,
			"tree":   ctx.Cates.GetTree(cats),
			"family": ctx.Cates.GetFamily(cats, gs.Cat_id),
			"user":   ctx.GetUser(),
			"cart":   ctx.GetCart(),
		}
		ctx.Render("./views/goods.html", data, lib.FuncMap)
	}
}
