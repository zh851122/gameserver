package main

import (
	Proto "awesomeProject/class3/src/classcode/Protocol"
	"awesomeProject/class3/src/classcode/Protocol/Proto2"
	"fmt"
	"github.com/golang/glog"
)

func (this *NetDataConn)HandleCletProtocol2Net(protocol2 interface{},ProtocolData map[string]interface{}){
	switch protocol2 {
	case float64(Proto2.Net_HeartBeatProto2):
		{
			//功能函数处理 -- 心跳
			this.HeartBeat(ProtocolData)
		}
	case float64(Proto2.Net_RelinkProto2):
		{
			//功能函数处理 -- 重新连接
			this.Relink(ProtocolData)
		}

	default:
		glog.Error("子协议：不存在！！！")

	}
	return
}
//重新连接
func (this *NetDataConn)Relink(ProtocolData map[string]interface{})  {
	//1 解析数据
	//2 update 网络数据
	//3 登录流程 ？？
	if ProtocolData["OpenID"] ==nil{
		panic("断线重连协议数据错误")
		return
	}
	StrOpenId := ProtocolData["OpenId"].(string)
	_=StrOpenId
	StrLoginName := ProtocolData["StrLoginName"].(string)
	StrLoginPW :=ProtocolData["StrLoginPW"].(string)
	//保存玩家数据
	playerdata := &NetDataConn{
		Connection: this.Connection,
		StrMd5:    (StrLoginName+StrLoginPW),
	}
	this.MapSafe.Put("PlayerUID"+"|connect",playerdata)
	//保存 ------
	G_PlayerData["123456"] = playerdata
	//服务器---->客户端
	data := &Proto2.Net_Relink{
		Proto:    Proto.GameNet_Proto,
		Proto2:    Proto2.Net_RelinkProto2,
		ISucc:true,
		OpenId:       "",
		StrLoginName: "",
		StrLoginPW:   "",
	}
	//发送数据给客户端了
	this.PlayerSendMessage(data)
	return
}
//心跳
func (this *NetDataConn)HeartBeat(ProtocolData map[string]interface{}){
	//1 解析协议数据
	//2 通过玩家的唯一id 去保存心跳数据 map[} data todo
	//3 timer ---超时踢人
	glog.Info("OpenId",ProtocolData["OpenId"])
	if ProtocolData["OpenId"] ==nil {
		panic("心跳协议数据错误")
		return
	}
	StrOpenId := ProtocolData["OpenId"].(string)
	fmt.Println("StrOpenId:",StrOpenId)
	if len(StrOpenId) == 65 {
		G_PlayeNet[StrOpenId]++
		if G_PlayeNet[StrOpenId]>100 {
			G_PlayeNet[StrOpenId] = 0
		}
		//返回数据
		//服务器-->客户端
		data := &Proto2.Net_HeartBeat{
			Proto:  Proto.GameData_Proto,
			Proto2: Proto2.Net_HeartBeatProto2,
			OpenId:    StrOpenId,
		}
		//发送数据给客户端
		this.PlayerSendMessage(data)
	}else {
		panic("心跳协议数据错误")
	}
	return

}