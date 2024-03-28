package serializer

import (
	"github.com/gin-gonic/gin"
)

type Error struct {
	Code     int
	Msg      string
	RawError error
}

func (e Error) Error() string {
	return e.Msg
}

func (e Error) WithError(err error) Error {
	e.RawError = err
	return e
}

var (
	CodeParamErr  = 400
	CodeServerErr = 500

	CodeEmailExist        = 10001
	CodeEncryptError      = 10002
	CodeInvalidSign       = 10003
	CodeSignTimeout       = 10004
	CodeNoPermission      = 10005
	CodeEmailSendErr      = 10006
	CodeCredentialInvalid = 10007
)

var (
	ServerError = Error{Code: CodeServerErr, Msg: "服务器内部错误"}
)

type ServerInternalError struct {
	Error
}

type ParamError struct {
	Error
}

type BizError struct {
	Error
}

func NewError(code int, msg string, err error) Error {
	return Error{
		Code:     code,
		Msg:      msg,
		RawError: err,
	}
}

func ParamErr(msg string, err error) Error {
	return NewError(CodeParamErr, msg, err)
}

func Err(code int, msg string, err error) Response {
	res := Response{
		Code: code,
		Msg:  msg,
	}

	if err != nil && gin.Mode() != gin.ReleaseMode {
		res.Error = err.Error()
	}
	return res
}
