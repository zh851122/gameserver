package Proto2
//GameNet_Proto 的子协议
const (
	INIT_PROTO2 = iota
	Net_HeartBeatProto2 // Net_HeartBeatProto2 ==1 心跳协议
	Net_Kicking_PlayerProto2 //Net_Kicking_PlayerProto2 == 2 踢人
	Net_RelinkProto2 // Net_RelinkProto2 == 3 断线重连协议
)
//断线重连
//玩家的结构协议 todo 玩家时长，在线时长 -----update
type Net_Relink struct {
	Proto int //主协议
	Proto2 int //子协议 -- 模块化的功能
	OpenId string //玩家唯一id  --server ---->client(多数不需要验证)
	StrLoginName string
	StrLoginPW string //加密的数据
	ISucc bool  //服务器返回的数据
}
// 踢人
type Net_Kicking_Player struct {
	Proto int //主协议
	Proto2 int //子协议 -- 模块化的功能
	OpenId string //玩家唯一id  --server ---->client(多数不需要验证)
	ErrorCode int //错误码 10001 10002
	StrMsg string //原因

}
//心跳
type Net_HeartBeat struct {
	Proto int //主协议
	Proto2 int //子协议 -- 模块化的功能
	OpenId string //玩家唯一id  --server ---->client(多数不需要验证)
}