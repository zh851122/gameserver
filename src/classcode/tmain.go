package main

import (
	"github.com/Golangltd/websocket_old/code.google.com/p/go.net/websocket"
	"github.com/golang/glog"
	"net/http"
	"os"
	"runtime"
)


func main()  {
	glog.Info("start-----111")
	strport :="8898"
	glog.Info(len(os.Args))
	if len(os.Args)>1{
		strport = os.Args[1]
	}

	glog.Info("游戏服务器启动开始---")
	glog.Info("本机几核：",runtime.NumCPU())
	glog.Info("当前端口:",strport)
	runtime.GOMAXPROCS(runtime.NumCPU())
	//http.HandleFunc("/",wwwGolangLtd)
	//websocket.Handler()
	http.Handle("/GolangLtd",websocket.Handler(wwwGolangLtd))
	http.ListenAndServe(":"+strport,nil)
	return
}