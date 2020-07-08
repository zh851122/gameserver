package main

import (
	"fmt"
	"github.com/Golangltd/websocket_old/code.google.com/p/go.net/websocket"
	"net/http"
	"runtime"
)

func init()  {

}
func main()  {
	fmt.Println("游戏服务器启动开始---")
	fmt.Println("本机几核：",runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())
	//http.HandleFunc("/",wwwGolangLtd)
	//websocket.Handler()
	http.Handle("/GolangLtd",websocket.Handler(wwwGolangLtd))
	http.ListenAndServe(":8888",nil)
	return
}