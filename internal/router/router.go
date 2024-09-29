package router

import (
	"StuService-Go/internal/handler"
	"StuService-Go/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	const pre = "/api"
	api := r.Group(pre)
	{
		uploadGroup := api.Group("/upload")
		{
			uploadGroup.Use(middleware.IsLogin)
			uploadGroup.POST("/picture", handler.UploadPicture)
			uploadGroup.POST("/multi_picture", handler.UploadMultiPicture)
		}
		userGroup := api.Group("/user")
		{
			// 未登陆
			userGroup.POST("/reg", handler.Register)
			userGroup.POST("/send_code", handler.SendCode)
			userGroup.POST("/login", handler.Login)

			// 已登陆
			userGroup.Use(middleware.IsLogin)
			userGroup.GET("/info", handler.GetUserInfo)
		}
	}
}
