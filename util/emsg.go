package emsg

const (
	Success = 10000
	Error   = 20000
	//用户错误
	UsernameExist     = 10001
	UsernameNoExist   = 10002
	PasswordWrong     = 10003
	NoLogin           = 10004
	SecretAnswerWrong = 10005
	//留言错误
	GetMessageFailed    = 20001
	PostMessageFailed   = 20002
	MessageNoExist      = 20003
	DeleteMessageFailed = 20004
)

var ErrorMsg map[int]string

//初始化状态码与错误信息对应的map
func init() {
	ErrorMsg = make(map[int]string)
	ErrorMsg[Success] = "操作成功"
	ErrorMsg[Error] = "操作失败"
	ErrorMsg[UsernameExist] = "用户名已存在"
	ErrorMsg[UsernameNoExist] = "用户名不存在"
	ErrorMsg[NoLogin] = "用户未登录"
	ErrorMsg[PasswordWrong] = "密码错误"
	ErrorMsg[SecretAnswerWrong] = "密保答案错误"
	ErrorMsg[GetMessageFailed] = "获取留言失败"
	ErrorMsg[PostMessageFailed] = "发布留言失败"
	ErrorMsg[MessageNoExist] = "留言不存在"
	ErrorMsg[DeleteMessageFailed] = "删除留言失败"
}

// @title 获取错误信息
// @description 接受状态码返回对应的错误信息
// @param code int "状态码"
// @return msg string "状态码对应的错误信息"
func GetErrorMsg(code int) string {
	return ErrorMsg[code]
}
