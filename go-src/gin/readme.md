对gin进行包裹，生成自定义的context上下文处理函数

```
g:=gin.Default()
eng:=New(g)
eng.Any("",func(ctx *Context){
    //ctx.User
    //ctx.Session
    //....
})
```