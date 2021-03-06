package Proto2

//GameDta_Proto 的字协议
const(
	INIt_PROTO2 = iota
	C2S_PlayerLoginProto2 // C2S_PlayerLoginProto2 == 1 客服端 ---->服务器 用户登录协议
	S2C_PlayerLoginProto2 //S2C_PlayerLoginProto2 == 2  服务器 -----> 客户端 用户登录协议

	C2S_ChooseRoomProto2 //C2S_ChooseRoomProto2 == 3 选择房间
	S2C_ChooseRoomProto2 //C2S_ChooseRoomProto2 == 4 选择房间
	C2S_PlayerRunProto2 //C2S_ChooseRoomProto2 == 5 模拟走路
	S2C_PlayerRunProto2 //C2S_ChooseRoomProto2 == 6 模拟走路
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
type C2S_PlayerRun struct {
	Head_Proto
	Itype int //1 登录，2 注册
	OpenId string // 微信
	StrRunX string
	StrRunY string
	StrRunZ string
}
//服务端------>客户端N 个客户端 ，(广播协议)
type S2C_PlayerRun struct {
	Head_Proto
	OpenId string // 微信
	StrRunX string
	StrRunY string
	StrRunZ string
	PlayerData *PlayerSt //玩家的结构
}
// 客户端---->服务端
type C2S_PlayerLogin struct {
	Head_Proto
	Itype int //1 登录，2 注册
	Code string // 微信 CODE
	StrLoginName string
	StrLoginPW string  //123456 ----> 还是加密数据？
	StrLoginEmail string //收取验证码
}
//服务端------>客户端
type S2C_PlayerLogin struct {
	Head_Proto
	PlayerData *PlayerSt //玩家的结构
}

