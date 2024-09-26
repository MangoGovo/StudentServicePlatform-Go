package handler

import (
	"StuService-Go/internal/apiException"
	"StuService-Go/internal/global"
	"StuService-Go/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func UploadPicture(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ServerError)
		return
	}

	// 保存图片
	picType := utils.GetFileType(file.Filename)
	dst := global.Config.GetString("file.imagePath") + "/" + utils.GetUUID() + picType

	if _, err := os.Stat(dst); err == nil {
		_ = c.AbortWithError(http.StatusOK, apiException.FileExistedError)
		return
	}

	if err := c.SaveUploadedFile(file, dst); err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ServerError)
		return
	}
	utils.JsonSuccess(c, gin.H{"picture_url": dst})
}
