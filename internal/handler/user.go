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
	utils.JsonSuccess(c, nil)
	return
	var data RegisterData
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.JsonFail(c, 200501, "参数错误，前端在干什么！乱传参数！")
		return
	}

	_, err := service.GetUserByUserName(data.Username)
	if err == nil {
		_ = c.AbortWithError(http.StatusOK, apiException.UserExistedError)
		return
	}
	//err = service.Register(&data)
	//if err != nil {
	//	_ = c.AbortWithError(http.StatusOK, apiException.ServerError)
	//	return
	//
	//}
	//utils.JsonSuccess(c, nil)

}
