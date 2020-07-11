package main

import (
	"flag"
	"fmt"
	"github.com/Golangltd/websocket_old/code.google.com/p/go.net/websocket"
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
var addr = flag.String("addr","127.0.0.1:8889","http service address")
func main()  {
	fmt.Println("Robot 客户端模拟------")
	url := "ws://" + *addr + "/GolangLtd"
	ws, err := websocket.Dial(url,"","test://golang/")
	if err !=nil{
		fmt.Println("err:",err.Error())
		return
	}
	_= ws

}
