package main

import (
	"flag"
	concurrent "github.com/Golangltd/go-concurrentMap"
)

//全局的网络信息结构
var G_PlayerData map[string]*NetDataConn
var G_PlayeNet map[string]int //心跳结构信息的存储结构
var G_PlayeNetSys map[string]int
var G_Net_Count map[string]int
var M *concurrent.ConcurrentMap //并发安全的
//初始化
func init()  {
	G_PlayerData = make(map[string]*NetDataConn)
	G_PlayeNet = make(map[string]int)
	G_PlayeNetSys = make(map[string]int)
	G_Net_Count = make(map[string]int)
	//并发安全的初始化
	M = concurrent.NewConcurrentMap()
	go G_timer()
	go G_timeout_kikc_Player()
	//命令的执行
	// server.exe -log_dir ="./"  -v=3
	//程序当中执行
	flag.Set("alsologtostderr","true")
	flag.Set("log_dir","./log")
	flag.Set("v","3")
	flag.Parse()
	return
}
