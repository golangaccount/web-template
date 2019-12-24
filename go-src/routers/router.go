package routers

import (
	"github.com/gin-gonic/gin"
	ugin "github.com/golangaccount/website-template/go-src/gin"
)

//RegistRouter 注册路由
func RegistRouter(g *gin.Engine) {
	engine := ugin.New(g)
	engine.Any("/test", func(ctx *ugin.Context) {
		ctx.Data(200, "", []byte("test"))
	})
}
