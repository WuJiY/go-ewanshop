package routes

import (
	"image/png"

	"github.com/haolifeng/go-ewanshop/framework"
	"github.com/haolifeng/go-ewanshop/lib/image"
)

func Capcha(ctx *framework.HandlerContext) {
	var r = image.MakeCapcha()

	ctx.SetSessionVal("__verify", r.Str)

	ctx.W.Header().Set("content-type", "image/png")
	err := png.Encode(ctx.W, r.Img)
	if err != nil {
		panic(err)
	}
}
