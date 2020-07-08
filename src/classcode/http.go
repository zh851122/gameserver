package main

import (
	"fmt"
	"github.com/Golangltd/websocket_old/code.google.com/p/go.net/websocket"
)

func wwwGolangLtd(ws *websocket.Conn)  {
	fmt.Println("golang 欢迎你",ws)
	data :=ws.Request().URL.Query().Get("data")
	fmt.Println("data:",data)
	//网络信息
	NetDataConntmp := &NetDataConn{
		Connection: ws,
		StrMd5:     "",
	}
	NetDataConntmp.PullFromClient()
}
