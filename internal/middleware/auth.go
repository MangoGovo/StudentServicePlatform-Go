package middleware

import (
	"StuService-Go/internal/apiException"
	"StuService-Go/internal/model"
	"StuService-Go/internal/service"
	"StuService-Go/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func IsLogin(c *gin.Context) {
	// 1. 从请求头中获取token
	tokenStr := c.Request.Header.Get("Authorization")
	if len(tokenStr) <= 7 {
		_ = c.AbortWithError(http.StatusOK, apiException.AuthExpired)
		return
	}
	tokenStr = tokenStr[7:]
	//utils.Log.Println(tokenStr)
	// 2. 解析token
	jwtUser, err := utils.ParseJwt(tokenStr)
	if err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.AuthExpired)
		return
	}

	// 3. 判断是否过期
	if time.Now().Unix() > jwtUser.ExpiresAt.Unix() {
		_ = c.AbortWithError(http.StatusOK, apiException.AuthExpired)
		return
	}

	// 4. 获取用户信息
	user, err := service.GetUserByID(jwtUser.UserID)
	if err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.UserNotExistError)
		return
	}
	c.Set("user", user)
	c.Set("isLogin", true)
}

func IsAdmin(c *gin.Context) {
	IsLogin(c)
	val, isExisted := c.Get("user")
	if !isExisted {
		return
	}
	user := val.(*model.User)
	if !(user.UserType == 1 || user.UserType == 2) {
		_ = c.AbortWithError(http.StatusOK, apiException.PermissionsNotAllowed)
		return
	}
}
