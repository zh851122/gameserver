package main

import (
	"github.com/Golangltd/websocket_old/code.google.com/p/go.net/websocket"
	"github.com/golang/glog"
)

func wwwGolangLtd(ws *websocket.Conn)  {
	glog.Info("golang 欢迎你",ws)
	data :=ws.Request().URL.Query().Get("data")
	glog.Info("data:"+data)
	//网络信息
	NetDataConntmp := &NetDataConn{
		Connection: ws,
		StrMd5:     "",
		MapSafe: M,
	}
	glog.Info("网络信息",NetDataConntmp)
	NetDataConntmp.PullFromClient()
}
