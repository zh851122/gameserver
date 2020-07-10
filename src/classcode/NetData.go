package main

import (
	Proto "awesomeProject/class3/src/classcode/Protocol"
	"awesomeProject/class3/src/classcode/Protocol/Proto2"
	"encoding/json"
	"github.com/Golangltd/websocket_old/code.google.com/p/go.net/websocket"
	"github.com/golang/glog"
	"reflect"
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
		glog.Error("err:"+err.Error())
		return nil,err
	}
	return result,nil
}

func (this *NetDataConn)SyncMeassgeFun(content string) {
	//1 字符串 -----》转换其他格式 必须高效 （大量并发情况下 依然不影响性能，游戏服务 计算密集型的）
	glog.Info(content)
	//2 已经通过第一步转换成我们所要的格式了，实现格式的处理函数：主协议、子协议、struct
	//3 处理函数实现
	var r Requestbody
	r.req = content
	if ProtocolData, err :=r.Json2Map();err==nil {
		//处理函数
		this.HandleCltProtocol(ProtocolData["Proto"],ProtocolData["Proto2"],ProtocolData)
	} else {
		glog.Error("解析失败："+err.Error())
	}

}
func typeof(v interface{}) string  {
	return reflect.TypeOf(v).String()
}
// 处理函数(底层函数了,必须面向所有的数据处理)
func (this *NetDataConn) HandleCltProtocol(protocol interface{},protocol2 interface{},ProtocolData map[string]interface{})  {
	//分发处理 -- 首先判断主协议存在，再判断子协议存在
	//fmt.Println(protocol)
	//fmt.Println(Proto.GameData_Proto)
	//
	//fmt.Println(typeof(protocol))
	//fmt.Println(typeof(Proto.GameData_Proto))
	switch protocol {
	case float64(Proto.GameData_Proto):
		{
			//子协议处理
			this.HandleCletProtocol2(protocol2,ProtocolData)
		}
	case float64(Proto.GameDataDB_Proto):
		{

		}
	default:
		glog.Error("主协议不存在！！！！")

	}

	return
}

//子协议的处理

func (this *NetDataConn)HandleCletProtocol2(protocol2 interface{},ProtocolData map[string]interface{}){
	switch protocol2 {
	case float64(Proto2.C2S_PlayerLoginProto2):
		{
		//功能函数处理 -- 用户登录协议
			this.PlayerLogin(ProtocolData)
		}
	default:
		glog.Error("子协议：不存在！！！")
	
	}
	return
}

//用户登录的协议

func (this *NetDataConn) PlayerLogin(ProtocolData map[string]interface{}) {
	//服务器的逻辑处理
	//获取用户发过来的数据信息
	/**
	1 StrLoginName
	2 StrLoginPW
	3 StrLoginEmail

	 */
	//1、获取client传过来的code
	//2、通过微信提供的接口--获取玩家的 == openid name 头像信息 ----
	//3、将用户信息保存到我们的数据库里（异步处理）
	// 检测用户数据是否存在，存在更新昵称，头像，否则保存
	//4、返回给客户端数据 ：== openid name
	if ProtocolData["StrLoginName"] == nil ||
		ProtocolData["StrLoginPW"] == nil ||
		ProtocolData["StrLoginEmail"] == nil {
		glog.Error(" 主协议 GameData_Proto,子协议 C2S_PlayerLoginProto2,登陆功能数据错误")
	}
	//玩家信息
	StrLoginName :=  ProtocolData["StrLoginName"].(string)
	StrLoginPW :=  ProtocolData["StrLoginPW"].(string)
	StrLoginEmail :=  ProtocolData["StrLoginEmail"].(string)
	glog.Info(StrLoginName+StrLoginEmail+StrLoginEmail+StrLoginPW)

	//数据库的保存 --发给dbserver
	//返回给客户端
	//channel 操作
	data:=&Proto2.S2C_PlayerLogin{
		Head_Proto: Proto2.Head_Proto{
			Proto:Proto.GameData_Proto,
			Proto2:Proto2.S2C_PlayerLoginProto2,
		},
		PlayerData: nil,
	}
	//发送数据给客户端
	this.PlayerSendMessage(data)
	return
}

// 发送给客户端的数据信息函数
//底层函数
func (this *NetDataConn)PlayerSendMessage(senddata interface{}) {
	// 1 消息序列化 interfac ->json
	b,errjson:=json.Marshal(senddata)
	if errjson!=nil{
		glog.Info(errjson.Error())
		return
	}
	//数据转换 json的byte数组--->string
	data := "data:"+string(b[0:len(b)])
	glog.Info(data)
	//发送客户端
	err:=websocket.JSON.Send(this.Connection,b)
	if err !=nil{
		glog.Info(err.Error())
		return
	}
	glog.Flush()
	return
}