package main

import (
	Proto "awesomeProject/class3/src/classcode/Protocol"
	"encoding/json"
	"fmt"
	"github.com/Golangltd/websocket_old/code.google.com/p/go.net/websocket"
)

/**
网络数据结构保存
1 websocket 的网络连接
2 StrMd5 房间的加密信息
 */
type NetDataConn struct {
	Connection *websocket.Conn
	StrMd5 string
}
// 结构体的 方法 接受者是指针类型的
func (this *NetDataConn) PullFromClient()  {
	//网络层处理，数据
	//1 针对服务器而言 一直等待消息的
	//for(){}
	for{
		var content string
		if err :=websocket.Message.Receive(this.Connection,&content); err !=nil{
			break
		}
		if len(content) == 0 {
			break
		}

		//go 并发编程使用
		go this.SyncMeassgeFun(content)
	}
	return
}
// 结构体数据类型
type Requestbody struct {
	req string
}
//json 转换为map:数据的处理
func (r *Requestbody) Json2Map() (s map[string]interface{},err error) {
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(r.req), &result); err != nil {
		fmt.Println("err:",err.Error())
		return nil,err
	}
	return result,nil
}

func (this *NetDataConn)SyncMeassgeFun(content string) {
	//1 字符串 -----》转换其他格式 必须高效 （大量并发情况下 依然不影响性能，游戏服务 计算密集型的）
	fmt.Println(content)
	//2 已经通过第一步转换成我们所要的格式了，实现格式的处理函数：主协议、子协议、struct
	//3 处理函数实现
	var r Requestbody
	r.req = content
	if ProtocolData, err :=r.Json2Map();err==nil {
		//处理函数
		this.HandleCltProtocol(ProtocolData["Proto"],ProtocolData["Proto2"],ProtocolData)
	} else {
		fmt.Println("解析失败：",err.Error())
	}

}

// 处理函数(底层函数了,必须面向所有的数据处理)
func (this *NetDataConn) HandleCltProtocol(protocol interface{},protocol2 interface{},ProtocolData map[string]interface{})  {
	//分发处理 -- 首先判断主协议存在，再判断子协议存在
	switch protocol {
	case Proto.GameData_Proto:
		{
			//子协议处理
		}
	case Proto.GameDataDB_Proto:
		{

		}
	default:
		panic("主协议不存在！！！！")

	}

	return
}