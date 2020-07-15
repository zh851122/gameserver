package main

import (
	Proto "awesomeProject/class3/src/classcode/Protocol"
	"awesomeProject/class3/src/classcode/Protocol/Proto2"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/Golangltd/websocket_old/code.google.com/p/go.net/websocket"
	"strings"
)

/**
1 模拟客户端发消息
2 客户端是并发创建 （并发访问服务器）
3 模拟真实客户登录
目的：
1 测试在线人数
2 框架的基础网络性能
3 框架的基础conn 存储
 */

/**
robot
1 模拟玩家的正常操作，例如 行走 跳跃 开枪 等等
2 做服务的性能测试，例如 并发量 内存 cpu 等等
3 压力测试
注意点：
1 模拟 ------> 多线程模拟 goroutine --server !!!!
首先
1 net 网络使用 websocket 进行连接
2 send 如何发送 ？？

 */
var addr = flag.String("addr","127.0.0.1:8898","http service address")
// 1 robot 客户端是可以一起链接的 ----> websocket.Dial 每次都返回一个
//2 多个 websocket.Dial ----> 多个客户端

func main()  {
	fmt.Println("Robot 客户端模拟------")
	url := "ws://" + *addr + "/GolangLtd"
	ws, err := websocket.Dial(url,"","test://golang/")
	if err !=nil{
		fmt.Println("err:",err.Error())
		return
	}
	go Send(ws,"Login")
	go Timer(ws)
	for{
		var content string
		err := websocket.Message.Receive(ws,&content)
		if err !=nil{
			fmt.Println("err:",err.Error())
			return
		}
		//decode
		fmt.Println(strings.Trim("","\""))
		fmt.Println(content)
		content = strings.Replace(content,"\"","",-1)
		contentstr,errr := base64Decode([]byte(content))
		if errr != nil {
			fmt.Println(errr)
		}
		fmt.Println("msg:",string(contentstr))
	}

}
func base64Decode(src []byte)([]byte,error)  {
	return base64.StdEncoding.DecodeString(string(src))
}
//消息流程
//1 针对消息结构进行数据的组装
//2 针对组装的数据进行一个数据格式的转换---->json
//3 json 数据之间发送到server

func Send(conn *websocket.Conn,Itype string)  {
	if Itype == "Login" {
		data :=&Proto2.C2S_PlayerLogin{
			Head_Proto:    Proto2.Head_Proto{
				Proto: Proto.GameData_Proto,
				Proto2: Proto2.C2S_PlayerLoginProto2,
			},
			Itype:         1,
			Code:          "",
			StrLoginName:  "user11",
			StrLoginPW:    "123456",
			StrLoginEmail: "dddd@qq.com",
		}
		//结构体转json
		//发送数据到服务器
		PlayerSendToServer(conn,data)
	}else if Itype == "HeartBeat" {

	}

	return
}

//公用send 函数
func PlayerSendToServer(conn *websocket.Conn,data interface{})  {
	//结构体转json
	jsons,err:=json.Marshal(data)
	if err !=nil{
		fmt.Println("err:",err.Error())
	}
	fmt.Println("jsons:",string(jsons))
	websocket.Message.Send(conn,jsons)
	return
}