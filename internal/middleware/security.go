package middleware

import (
	"StuService-Go/internal/apiException"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// Security 防压测中间件
func Security() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 彭海林你喜欢用 Apifox 是吧
		UserAgent := c.GetHeader("User-Agent")
		if strings.Contains(UserAgent, "Apifox") {
			_ = c.AbortWithError(http.StatusOK, apiException.Phlin)
			return
		}
	}
}
