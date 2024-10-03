package middleware

import (
	"github.com/gin-gonic/gin"
)

// Security 防压测中间件
func Security() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 彭海林你喜欢用 Apifox 是吧
		//UserAgent := c.GetHeader("User-Agent")
		//if strings.Contains(UserAgent, "Apifox") {
		//	utils.Log.Printf("[%s]使用apifox发起了请求", c.ClientIP())
		//	_ = c.AbortWithError(http.StatusOK, apiException.Phlin)
		//	return
		//}
	}
}
