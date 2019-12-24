package gin

import (
	"github.com/gin-gonic/gin"
)

//Context 运行上下文
type Context struct {
	*gin.Context
}

//HandlerFunc 处理函数
type HandlerFunc func(*Context)
