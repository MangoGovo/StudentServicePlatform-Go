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
			userGroup.POST("/edit_info", handler.EditUserInfo)
		}
		stuGroup := api.Group("/student")
		{
			stuGroup.Use(middleware.IsLogin)
			stuGroup.POST("/feedback", handler.PostFeedback)
			stuGroup.GET("/feedback", handler.QueryFeedback)
			stuGroup.DELETE("/feedback", handler.DeleteFeedback)
			stuGroup.PUT("/feedback", handler.UpdateFeedback)
			stuGroup.GET("/feedback_list", handler.GetFeedbackList)

		}
		//adminGroup := api.Group("/admin"){}
		//sudoGroup := api.Group("/sudo"){}
	}
}
