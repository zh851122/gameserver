package main

import (
	Proto "awesomeProject/class3/src/classcode/Protocol"
	"awesomeProject/class3/src/classcode/Protocol/Proto2"
	"github.com/golang/glog"
	"strings"
	"time"
)

func Strings_Split(Data string,Split string)[]string  {
	return strings.Split(Data,Split)
}
// 超时踢人
func G_timeout_kikc_Player()  {
	for  {
		select {
		case <-time.After(10*time.Second):
			{
			 //1 获取心跳数据 -- 玩家的 测试一个玩家 data[A] =1
			// 2 玩家的心跳保存下来 -- 临时保存 datatemp[A] = 1
			// 3 每10s 对比一次：临时的保存数据 与 我们的心跳数据是否相同 data[A] == datatmp[A]
			// 4 30s 还是没有变化 kick player A
			// 并发安全map 优化
			for itr:=M.Iterator();itr.HasNext();{
				k,v, _ :=itr.Next()
				//取分隔符
				strsplit := Strings_Split(k.(string),"|")
				for i :=0;i<len(strsplit);i++ {
					if len(strsplit) < 2 {
						continue
					}
					//进行数据的查询类型
					switch v.(interface{}).(type) {
					case *NetDataConn:
						{
						//判断 链接是不是 connect
						if  k == "connect"{
							data := &Proto2.Net_Kicking_Player{
								Proto:  Proto.GameNet_Proto,
								Proto2: Proto2.Net_Kicking_PlayerProto2,
								ErrorCode: 10001,
							}
							v.(interface{}).(*NetDataConn).PlayerSendMessage(data)
							}
						}

						}
					}
				}
			}
			//if G_Net_Count["12345"] >=3 {
			//	//踢人
			//	data := &Proto2.Net_Kicking_Player{
			//		Protocol:  Proto.GameNet_Proto,
			//		Protocol2: Proto2.Net_Kicking_PlayerProto2,
			//		ErrorCode:  10001,
			//	}
			//	G_PlayerData["12345"].PlayerSendMessage(data)
			//	//关闭连接
			//	G_PlayerData["12345"].Connection.Close()
			//	G_Net_Count["12345"] = 0
			//	continue
			//}
			//if len(G_PlayeNetSys) == 0 {
			//	G_PlayeNetSys["12345"] = G_PlayeNet["12345"]
			//}else {
			//	if G_PlayeNetSys["12345"] == G_PlayeNet["12345"]{
			//		G_Net_Count["12345"] ++
			//	}
			//}

			}

		}

}

//数据推送给客户端定时
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

//广播函数处理

func Broadcast(data interface{})  {
	for itr:=M.Iterator();itr.HasNext();{
		k,v, _ :=itr.Next()
		//取分隔符
		glog.Info("广播start ----")
		glog.Info(k)
		glog.Info(v)
		strsplit := Strings_Split(k.(string),"|")
		for i :=0;i<len(strsplit);i++ {
			if len(strsplit) < 2 {
				continue
			}
			//进行数据的查询类型
			switch v.(interface{}).(type) {
			case *NetDataConn:
				{
					//判断 链接是不是 connect
					if  k == "connect"{
						v.(interface{}).(*NetDataConn).PlayerSendMessage(data)
					}
				}

			}
		}
	}
}