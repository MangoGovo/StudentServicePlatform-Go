package handler

import (
	"StuService-Go/internal/apiException"
	"StuService-Go/internal/service"
	"StuService-Go/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterData struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`  //昵称
	Password string `json:"-"`         //密码
	UserType int    `json:"user_type"` //用户类型
}

func Register(c *gin.Context) {
	var data RegisterData

	if err := c.ShouldBindJSON(&data); err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ParamsError)
		return
	}
	email := data.Username
	// 校验验证码
	if !utils.IsValidEmail(email) {
		_ = c.AbortWithError(http.StatusOK, apiException.ParamsError)
		return
	}
	if _, err := service.GetUserByUserName(data.Username); err == nil {
		_ = c.AbortWithError(http.StatusOK, apiException.UserExistedError)
		return
	}
	verifyCode := utils.GenerateVerifyCode(6)
	if err := service.SendVerifyCode(email, verifyCode); err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.SendVerifyCodeError)
	}
	return
}
