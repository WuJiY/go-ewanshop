package user

import (
	"github.com/haolifeng/go-testShop/framework"
)

func Logout(ctx *framework.HandlerContext) {
	ctx.SetUser(nil)
	ctx.Redirect("./login.go")
}
