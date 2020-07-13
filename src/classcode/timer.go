package main

import (
	Proto "awesomeProject/class3/src/classcode/Protocol"
	"awesomeProject/class3/src/classcode/Protocol/Proto2"
	"github.com/golang/glog"
	"time"
)

func G_timer()  {
	for  {
		select {
		case <-time.After(20*time.Second):
			{
			if len(G_PlayerData) == 0 {
				continue
			}
			if G_PlayerData["123456"] != nil{
				data:=&Proto2.S2C_PlayerLogin{
					Head_Proto: Proto2.Head_Proto{
						Proto:Proto.GameData_Proto,
						Proto2:Proto2.S2C_PlayerLoginProto2,
					},
					PlayerData: nil,
				}
				G_PlayerData["123456"].PlayerSendMessage(data)
				glog.Info("data:",data)
			}

			}

		}
	}
}
