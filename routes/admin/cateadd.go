package admin

import (
	"github.com/haolifeng/go-ewanshop/framework"
	"github.com/haolifeng/go-ewanshop/lib"
	"github.com/haolifeng/go-ewanshop/models"
	"time"
)

func CateAdd(ctx *framework.HandlerContext) {
	var isAdmin = false
	if v, ok := ctx.GetSessionVal("isAdmin"); ok {
		isAdmin = v.(bool)
	}
	if !isAdmin {
		ctx.Redirect("/routes/admin/login.go")
	}

	if ctx.R.Method == "GET" {
		var allData = ctx.Cates.Find()
		var data = map[string]interface{}{
			"tree": ctx.Cates.GetTree(allData),
		}
		ctx.Render("./views/admin/cateadd.html", data, lib.FuncMap)
		return
	}

	if ctx.R.Method == "POST" {
		var data = map[string]interface{}{
			"cat_name":  ctx.R.FormValue("cat_name"),
			"intro":     ctx.R.FormValue("cat_desc"),
			"parent_id": ctx.R.FormValue("parent_id"),
			"oid":       models.MakeOid(),
			"createAt": time.Now(),
		}
		ctx.Cates.Add(data)
		ctx.Redirect("./catelist.go")
	}
}
