package gin

import (
	"github.com/gin-gonic/gin"
)

type engine struct {
	*gin.Engine
	*RouterGroup
}

//DefineContextKey 自定义context key
const DefineContextKey = "_golangaccount/context"

//New 生成一个转换后的engine
func New(g *gin.Engine) *engine {

	g.Use(func(ctx *gin.Context) {
		ctx.Set(DefineContextKey, &Context{Context: ctx})
	})
	return &engine{
		Engine: g,
		RouterGroup: &RouterGroup{
			RouterGroup: &g.RouterGroup,
		},
	}
}
