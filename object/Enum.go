package object

type ErrTypeCode int

const (
	ErrTypeCodeNoError            ErrTypeCode = 200
	ErrTypeCodeTokenTimeout       ErrTypeCode = 10000
	ErrTypeCodeLoginFailed        ErrTypeCode = 10001
	ErrTypeCodeExceedMaxSearchDay ErrTypeCode = 10002
)

type ErrTypeMsg string

const (
	ErrTypeMsgNoError            ErrTypeMsg = "success"
	ErrTypeMsgTokenTimeout       ErrTypeMsg = "token is invalid"
	ErrTypeMsgLoginFailed        ErrTypeMsg = "login failed"
	ErrTypeMsgExceedMaxSearchDay ErrTypeMsg = "查询间隔不能大于%d天"
)
