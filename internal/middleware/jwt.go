package middleware

import (
	"JH_2024_MJJ/internal/service"
	"JH_2024_MJJ/pkg/utils"
	"github.com/gin-gonic/gin"
)

func IsLogin(c *gin.Context) {
	isLogin := service.IsLogin(c.GetHeader("token"))
	if !isLogin {
		utils.JsonResponse(c, 200, 200404, "登录过期", nil)
		c.Abort()
	}
}

func IsAdmin(c *gin.Context) {
	isAdmin := service.IsAdmin(c.GetHeader("token"))
	if !isAdmin {
		utils.JsonResponse(c, 200, 200404, "权限不足", nil)
		c.Abort()
	}
}
