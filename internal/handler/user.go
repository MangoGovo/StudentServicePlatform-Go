package handler

import (
	"StuService-Go/internal/apiException"
	"StuService-Go/internal/model"
	"StuService-Go/internal/service"
	"StuService-Go/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SendCodeData struct {
	Email string `json:"email"`
}

func SendCode(c *gin.Context) {
	var data SendCodeData
	if err := c.ShouldBindJSON(&data); err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ParamsError)
		return
	}
	email := data.Email
	if !utils.IsValidEmail(email) {
		_ = c.AbortWithError(http.StatusOK, apiException.ParamsError)
		return
	}
	verifyCode := utils.GenerateVerifyCode(6)
	if err := service.SendVerifyCode(email, verifyCode); err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.SendVerifyCodeError)
		return
	}
	utils.JsonSuccess(c, nil)
}

type RegisterData struct {
	Username   string `json:"username"`
	Nickname   string `json:"nickname"`  //昵称
	Password   string `json:"password"`  //密码
	UserType   int    `json:"user_type"` //用户类型
	VerifyCode string `json:"verify_code"`
}

func Register(c *gin.Context) {
	var data RegisterData

	if err := c.ShouldBindJSON(&data); err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ParamsError)
		return
	}
	email := data.Username

	// 校验邮箱是否合法
	if !utils.IsValidEmail(email) {
		_ = c.AbortWithError(http.StatusOK, apiException.ParamsError)
		return
	}

	// 校验密码长度
	pwdLen := len(data.Password)
	if pwdLen < 8 || pwdLen > 16 {
		_ = c.AbortWithError(http.StatusOK, apiException.ParamsError)
		return
	}

	// 校验用户是否重复
	if _, err := service.GetUserByUserName(data.Username); err == nil {
		_ = c.AbortWithError(http.StatusOK, apiException.UserExistedError)
		return
	}

	// 校验验证码
	utils.Log.Println(email)
	cachedVerifyCode, err := service.GetVerifyCode(email)
	if err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.VerifyCodeExpired)
		return
	}

	if data.VerifyCode != cachedVerifyCode {
		_ = c.AbortWithError(http.StatusOK, apiException.VerifyCodeError)
		return
	}

	err = service.Register(data.Username, data.Nickname, data.Password, data.UserType)
	if err != nil {
		_ = c.AbortWithError(http.StatusOK, err)
		return
	}

	utils.JsonSuccess(c, nil)
}

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"` // MD5 Encrypted
}

func Login(c *gin.Context) {
	var data LoginData
	if err := c.ShouldBindJSON(&data); err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ParamsError)
		return
	}

	// 1. 判断密码是否为MD5加密
	if !utils.CheckMD5(data.Password) {
		_ = c.AbortWithError(http.StatusOK, apiException.ParamsError)
		return
	}

	// 2. 判断用户是否存在
	user, err := service.GetUserByUserName(data.Username)
	if err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.UserNotExistError)
		return
	}

	//	3. 密码校验
	if user.Password != data.Password {
		_ = c.AbortWithError(http.StatusOK, apiException.PwdWrongError)
		return
	}

	// 4. 生成Token
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ServerError)
		return
	}

	// 5. 获取UserType
	utils.JsonSuccess(c, gin.H{
		"token":     token,
		"user_type": user.UserType,
	})
}

func GetUserInfo(c *gin.Context) {
	user := c.MustGet("user").(*model.User)
	utils.JsonSuccess(c, user)
}

type editUserInfoData struct {
	Nickname     string `json:"nickname"`
	Gender       int    `json:"gender"`
	Introduction string `json:"introduction"`
}

func EditUserInfo(c *gin.Context) {
	var data editUserInfoData
	if err := c.ShouldBindJSON(&data); err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ParamsError)
		return
	}
	user := c.MustGet("user").(*model.User)

	if err := service.UpdateUser(&model.User{
		ID:           user.ID,
		Username:     user.Username,
		Nickname:     data.Nickname,
		Password:     user.Password,
		Gender:       data.Gender,
		Introduction: data.Introduction,
		UserType:     user.UserType,
	}); err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ParamsError)
		return
	}
	utils.JsonSuccess(c, nil)
}
