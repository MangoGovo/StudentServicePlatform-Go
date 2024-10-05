package apiException

import "net/http"

type Error struct {
	StatusCode int    `json:"-"`
	Code       int    `json:"code"`
	Msg        string `json:"msg"`
}

var (
	ServerError           = NewError(http.StatusInternalServerError, 200500, "系统异常，请稍后重试!")
	UserExistedError      = NewError(http.StatusInternalServerError, 200501, "用户已经存在")
	FileExistedError      = NewError(http.StatusInternalServerError, 200502, "文件已经存在")
	ParamsError           = NewError(http.StatusInternalServerError, 200503, "参数错误")
	FileTypeError         = NewError(http.StatusInternalServerError, 200504, "文件类型错误")
	SendVerifyCodeError   = NewError(http.StatusInternalServerError, 200505, "验证码发送失败")
	PermissionsNotAllowed = NewError(http.StatusInternalServerError, 200506, "权限不足")
	VerifyCodeError       = NewError(http.StatusInternalServerError, 200507, "验证码错误")
	VerifyCodeExpired     = NewError(http.StatusInternalServerError, 200508, "验证码已过期，请重新发送")
	UserNotExistError     = NewError(http.StatusInternalServerError, 200509, "用户不存在")
	PwdWrongError         = NewError(http.StatusInternalServerError, 200510, "密码错误")
	AuthExpired           = NewError(http.StatusInternalServerError, 200511, "登陆状态已过期，请重新登陆")
	FeedbackNotHandled    = NewError(http.StatusInternalServerError, 200512, "问题反馈仍未被处理,请待处理后再打分")
	LimitExceeded         = NewError(http.StatusInternalServerError, 200513, "请求频率过快请稍后再试")

	Phlin = NewError(http.StatusInternalServerError, 114514, "我猜你是彭海林")
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
