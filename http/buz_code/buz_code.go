package buz_code

type Code uint8

const (
	CODE_OK                       Code = iota //ok 0
	CODE_INVALID_ARGS                         //参数错误 1
	CODE_SERVER_ERROR                         //服务器内部错误 2
	CODE_ENTERPRISE_CREATE_FAILED             //新建企业失败 3
	CODE_ENTERPRISE_UPDATE_FAILED             //更新企业失败 4
)
