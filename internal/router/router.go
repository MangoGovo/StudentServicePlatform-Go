package router

import (
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	const pre = "/api"
	api := r.Group(pre)
	{
		userGroup := api.Group("/user")
		{
			userGroup.POST("register", user.Register)
			//userGroup.POST("login", user.Login)
		}
	}
}
