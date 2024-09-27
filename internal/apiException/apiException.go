package apiException

import "net/http"

type Error struct {
	StatusCode int    `json:"-"`
	Code       int    `json:"code"`
	Msg        string `json:"msg"`
}

var (
	ServerError      = NewError(http.StatusInternalServerError, 200500, "系统异常，请稍后重试!")
	UserExistedError = NewError(http.StatusInternalServerError, 200501, "用户已经存在")
	FileExistedError = NewError(http.StatusInternalServerError, 200502, "文件已经存在")
	ParamError       = NewError(http.StatusInternalServerError, 200503, "参数错误")
	FileTypeError    = NewError(http.StatusInternalServerError, 200504, "文件类型错误")
)

func OtherError(message string) *Error {
	return NewError(http.StatusForbidden, 100403, message)
}

func (e *Error) Error() string {
	return e.Msg
}

func NewError(statusCode, Code int, msg string) *Error {
	return &Error{
		StatusCode: statusCode,
		Code:       Code,
		Msg:        msg,
	}
}
