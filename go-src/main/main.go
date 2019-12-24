package main

import (
	"flag"
	"net/http"
	"time"

	"github.com/facebookgo/httpdown"
	"github.com/gin-gonic/gin"
	"github.com/golangaccount/website-template/go-src/routers"
)

//热更新，关闭程序（等所有的请求全部退出后程序退出），启动新程序。原有的请求继续处理，新的请求新启动的程序进行处理
//rename file
//spc file
//chmod +x
//kill
//start
func main() {
	g := gin.Default()
	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: g,
	}

	routers.RegistRouter(g)

	hd := &httpdown.HTTP{
		StopTimeout: 10 * time.Second,
		KillTimeout: 1 * time.Second,
	}

	flag.StringVar(&server.Addr, "addr", server.Addr, "http address")
	flag.DurationVar(&hd.StopTimeout, "stop-timeout", hd.StopTimeout, "stop timeout")
	flag.DurationVar(&hd.KillTimeout, "kill-timeout", hd.KillTimeout, "kill timeout")
	flag.Parse()

	if err := httpdown.ListenAndServe(server, hd); err != nil {
		panic(err)
	}
}
