package router

import (
	"StuService-Go/internal/handler"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	const pre = "/api"
	api := r.Group(pre)
	{
		uploadGroup := api.Group("/upload")
		{
			uploadGroup.POST("/picture", handler.UploadPicture)
			uploadGroup.POST("/multi_picture", handler.UploadMultiPicture)
		}
		userGroup := api.Group("/user")
		{
			userGroup.POST("reg", handler.Register)
			//userGroup.POST("login", user.Login)
		}
	}
}
