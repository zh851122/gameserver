package main

import (
	Proto "awesomeProject/class3/src/classcode/Protocol"
	"awesomeProject/class3/src/classcode/Protocol/Proto2"
	"github.com/Golangltd/websocket_old/code.google.com/p/go.net/websocket"

	"time"
)

//定时发心跳
func Timer(conn *websocket.Conn)  {
	itimer := time.NewTicker(1 * time.Second)
	for  {
		select {
		case <-itimer.C:
			//心跳
			//1 组装
			StrOpenId := "12345678901234567890123456789012345678901234567890123456789012345"
			data := &Proto2.Net_HeartBeat{
				Proto:  Proto.GameNet_Proto,
				Proto2: Proto2.Net_HeartBeatProto2,
				OpenId:  StrOpenId ,
			}
			//发送数据到服务器
			PlayerSendToServer(conn,data)
			// --------------
			dataro := &Proto2.S2C_PlayerRun{
				Head_Proto: Proto2.Head_Proto{
					Proto:Proto.GameData_Proto,
					Proto2:Proto2.C2S_PlayerRunProto2,
				},
				OpenId:     StrOpenId,
				StrRunX:    "22",
				StrRunY:    "22",
				StrRunZ:    "22",
			}
			PlayerSendToServer(conn,dataro)
			//run 的消息
		}
		
	}
}
