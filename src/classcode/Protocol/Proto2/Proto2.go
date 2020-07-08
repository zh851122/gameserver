package Proto2

//GameDta_Proto 的字协议
const(
	INIt_PROTO2 = iota
	C2S_PlayerLoginProto2 // C2S_PlayerLoginProto2 == 1 客服端 ---->服务器 用户登录协议
	S2C_PlayerLoginProto2 //S2C_PlayerLoginProto2 == 2  服务器 -----> 客户端 用户登录协议

	C2S_ChooseRoomProto2 //C2S_ChooseRoomProto2 == 3 选择房间
	S2C_ChooseRoomProto2 //C2S_ChooseRoomProto2 == 4 选择房间
)

type PlayerSt struct {
	UID int
	PlayerName string
	OpenID string
}
// 功能结构
type Head_Proto struct {
	Proto int // 主协议 -- 模块化
	Proto2 int // 子协议 -- 模块化的功能
}
//----------------------------------------------------------------------------------------------------------------
// 客户端---->服务端
type C2S_PlayerLogin struct {
	Head_Proto
	Code string // 微信 CODE
}
//服务端------>客户端
type S2C_PlayerLogin struct {
	Head_Proto
	PlayerData *PlayerSt //玩家的结构
}

