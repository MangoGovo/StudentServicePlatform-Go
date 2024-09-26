package service

// 管理员密码错误
type adminPasswordIncorrectError struct{}

func (e *adminPasswordIncorrectError) Error() string {
	return "管理员密码错误，注册失败"
}

// 账号密码不合法
type invalidUsernameOrPasswordError struct{}

func (e *invalidUsernameOrPasswordError) Error() string {
	return "账号密码不合法，注册失败"
}

// 用户类型不合法
type invalidUserTypeError struct{}

func (e *invalidUserTypeError) Error() string {
	return "用户类型不合法，注册失败"
}
