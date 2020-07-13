package main

import (
	"flag"
	"github.com/Golangltd/websocket_old/code.google.com/p/go.net/websocket"
	"github.com/golang/glog"
	"net/http"
	"os"
	"runtime"
)
//全局的网络信息结构
var G_PlayerData map[string]*NetDataConn
//初始化
func init()  {
	G_PlayerData = make(map[string]*NetDataConn)
	go G_timer()
	//命令的执行
	// server.exe -log_dir ="./"  -v=3
	//程序当中执行
	flag.Set("alsologtostderr","true")
	flag.Set("log_dir","./log")
	flag.Set("v","3")
	flag.Parse()
	return
}

func main()  {
	glog.Info("start-----111")
	strport :="8889"
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
	http.ListenAndServe(":8898",nil)
	return
}