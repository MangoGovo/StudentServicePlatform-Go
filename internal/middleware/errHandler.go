package middleware

import (
	"StuService-Go/internal/apiException"
	"StuService-Go/pkg/utils"
	"errors"
	"github.com/gin-gonic/gin"
)

func ErrHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			e := c.Errors.Last()
			err := e.Err
			if err != nil {
				var apiErr *apiException.Error
				var e *apiException.Error
				if errors.As(err, &e) {
					apiErr = e
				}
				utils.Log.Printf("[ip:%s]%s", c.ClientIP(), apiErr.Msg)
				utils.JsonFail(c, apiErr.Code, apiErr.Msg)
				return
			}
		}
	}
}
