package cors

import (
	"github.com/kataras/iris/v12"
	"tonovel/bootstrap"
)

func New() iris.Handler {
	return func(ctx iris.Context) {
		ctx.Header("Access-Control-Allow-Origin", ctx.GetHeader("origin"))
		ctx.Header("Access-Control-Allow-Credentials", "true")
		ctx.Header("Access-Control-Allow-methods", "GET, POST, OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Origin,Content-Type")
		ctx.Next()
	}
}

func Configure(b *bootstrap.Bootstrapper) {
	h := New()
	b.UseGlobal(h)
}