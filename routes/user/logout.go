package user

import (
	"github.com/haolifeng/go-ewanshop/framework"
)

func Logout(ctx *framework.HandlerContext) {
	ctx.SetUser(nil)
	ctx.Redirect("./login.go")
}
