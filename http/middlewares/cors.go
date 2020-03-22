package middlewares

import (
	"github.com/kataras/iris/v12"
)

var cors  = &corsMiddleware{}

type corsMiddleware struct {}
func (m *corsMiddleware) New() iris.Handler {
	return func(ctx iris.Context) {
		ctx.Header("Access-Control-Allow-Origin", ctx.GetHeader("origin"))
		ctx.Header("Access-Control-Allow-Credentials", "true")
		ctx.Header("Access-Control-Allow-methods", "GET, POST, OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Origin,Content-Type")
		ctx.Next()
	}
}