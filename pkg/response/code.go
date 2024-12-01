package res

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeServerBusy
	CodeFail
	CodeAuthFail
	CodeTokenInvalid
)

var codeMsg = map[ResCode]interface{}{
	CodeSuccess:      "success",
	CodeInvalidParam: "参数错误",
	CodeServerBusy:   "服务器繁忙",
	CodeFail:         "失败",
	CodeAuthFail:     "鉴权失败",
	CodeTokenInvalid: "token无效",
}

func (c ResCode) Msg() interface{} {
	msg, ok := codeMsg[c]
	if !ok {
		msg = codeMsg[CodeServerBusy]
	}

	return msg
}
