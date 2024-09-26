package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func JsonResponse(c *gin.Context, httpStatus int, code int, msg string, data interface{}) {
	c.JSON(httpStatus, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	},
	)
}

func JsonSuccess(c *gin.Context, data interface{}) {
	JsonResponse(c, http.StatusOK, 200, "success", data)
}

func JsonFail(c *gin.Context, code int, msg string) {
	JsonResponse(c, http.StatusOK, code, msg, nil)
}
