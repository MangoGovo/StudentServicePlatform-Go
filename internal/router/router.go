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
			stuGroup.POST("/rate", handler.RateFeedback)
			stuGroup.POST("/comment", handler.CommentFeedback)
			stuGroup.DELETE("/comment", handler.DeleteComment)
		}

		adminGroup := api.Group("/admin")
		{
			adminGroup.Use(middleware.IsAdmin)
			adminGroup.POST("/order", handler.Order)
			adminGroup.POST("/undo_order", handler.UndoOrder)
			adminGroup.POST("/rubbish", handler.Rubbish)
			adminGroup.POST("/comment", handler.AdminComment)
			adminGroup.DELETE("/comment", handler.AdminDelComment)
			adminGroup.GET("/feedback_list", handler.AdminGetFeedbackList)
			adminGroup.GET("/feedback", handler.AdminQueryFeedback)
		}
		//sudoGroup := api.Group("/sudo"){}
	}
}
