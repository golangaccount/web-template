package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	//IRoutes router
	IRoutes interface {
		Use(...HandlerFunc) IRoutes

		Handle(string, string, ...HandlerFunc) IRoutes
		Any(string, ...HandlerFunc) IRoutes
		GET(string, ...HandlerFunc) IRoutes
		POST(string, ...HandlerFunc) IRoutes
		DELETE(string, ...HandlerFunc) IRoutes
		PATCH(string, ...HandlerFunc) IRoutes
		PUT(string, ...HandlerFunc) IRoutes
		OPTIONS(string, ...HandlerFunc) IRoutes
		HEAD(string, ...HandlerFunc) IRoutes

		StaticFile(string, string) IRoutes
		Static(string, string) IRoutes
		StaticFS(string, http.FileSystem) IRoutes
	}
	//RouterGroup 分组路由
	RouterGroup struct {
		*gin.RouterGroup
	}
)

func newHandlerFunc(f HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		v, _ := ctx.Get(DefineContextKey)
		f(v.(*Context))
	}
}

func newHandlersChain(fs []HandlerFunc) gin.HandlersChain {
	handler := make([]gin.HandlerFunc, len(fs))
	for i, item := range fs {
		handler[i] = newHandlerFunc(item)
	}
	return handler
}

//Any 所有的请求方式
func (group *RouterGroup) Any(relativePath string, handlers ...HandlerFunc) IRoutes {
	group.RouterGroup.Any(relativePath, newHandlersChain(handlers)...)
	return group
}

//DELETE method=delete
func (group *RouterGroup) DELETE(relativePath string, handlers ...HandlerFunc) IRoutes {
	group.RouterGroup.DELETE(relativePath, newHandlersChain(handlers)...)
	return group
}

//GET GET
func (group *RouterGroup) GET(relativePath string, handlers ...HandlerFunc) IRoutes {
	group.RouterGroup.GET(relativePath, newHandlersChain(handlers)...)
	return group
}

//Group Group
func (group *RouterGroup) Group(relativePath string, handlers ...HandlerFunc) *RouterGroup {
	return &RouterGroup{
		RouterGroup: group.RouterGroup.Group(relativePath, newHandlersChain(handlers)...),
	}
}

//HEAD HEAD
func (group *RouterGroup) HEAD(relativePath string, handlers ...HandlerFunc) IRoutes {
	group.RouterGroup.HEAD(relativePath, newHandlersChain(handlers)...)
	return group
}

//Handle Handle
func (group *RouterGroup) Handle(httpMethod, relativePath string, handlers ...HandlerFunc) IRoutes {
	group.RouterGroup.Handle(httpMethod, relativePath, newHandlersChain(handlers)...)
	return group
}

//OPTIONS OPTIONS
func (group *RouterGroup) OPTIONS(relativePath string, handlers ...HandlerFunc) IRoutes {
	group.RouterGroup.OPTIONS(relativePath, newHandlersChain(handlers)...)
	return group
}

//PATCH PATCH
func (group *RouterGroup) PATCH(relativePath string, handlers ...HandlerFunc) IRoutes {
	group.RouterGroup.PATCH(relativePath, newHandlersChain(handlers)...)
	return group
}

//POST POST
func (group *RouterGroup) POST(relativePath string, handlers ...HandlerFunc) IRoutes {
	group.RouterGroup.POST(relativePath, newHandlersChain(handlers)...)
	return group
}

//PUT PUT
func (group *RouterGroup) PUT(relativePath string, handlers ...HandlerFunc) IRoutes {
	group.RouterGroup.PUT(relativePath, newHandlersChain(handlers)...)
	return group
}

//Static Static
func (group *RouterGroup) Static(relativePath, root string) IRoutes {
	group.Static(relativePath, root)
	return group
}

//StaticFS StaticFS
func (group *RouterGroup) StaticFS(relativePath string, fs http.FileSystem) IRoutes {
	group.StaticFS(relativePath, fs)
	return group
}

//StaticFile StaticFile
func (group *RouterGroup) StaticFile(relativePath, filepath string) IRoutes {
	group.StaticFile(relativePath, filepath)
	return group
}

//Use Use
func (group *RouterGroup) Use(middleware ...HandlerFunc) IRoutes {
	group.RouterGroup.Use(newHandlersChain(middleware)...)
	return group
}
