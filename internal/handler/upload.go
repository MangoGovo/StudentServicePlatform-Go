package handler

import (
	"StuService-Go/internal/apiException"
	"StuService-Go/internal/global"
	"StuService-Go/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func isImageFile(filename string) bool {
	//	判断一个文件是否是图片
	imageExtension := global.Config.GetStringSlice("file.imageExtension")
	for _, ext := range imageExtension {
		if ext == filename {
			return true
		}
	}
	return false
}
func UploadPicture(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ServerError)
		return
	}

	// 保存图片
	picType := utils.GetFileType(file.Filename)
	if !isImageFile(picType) {
		_ = c.AbortWithError(http.StatusOK, apiException.FileTypeError)
		return
	}

	dst := global.Config.GetString("file.imagePath") + "/" + utils.GetUUID() + picType

	if err := c.SaveUploadedFile(file, dst); err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ServerError)
		return
	}
	utils.JsonSuccess(c, gin.H{"picture_url": dst})
}

func UploadMultiPicture(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ServerError)
		return
	}
	files := form.File["image"]
	PictureList := make([]string, len(files))

	// 保存图片
	for index, file := range files {
		picType := utils.GetFileType(file.Filename)
		if !isImageFile(picType) {
			_ = c.AbortWithError(http.StatusOK, apiException.FileTypeError)
			return
		}
		dst := global.Config.GetString("file.imagePath") + "/" + utils.GetUUID() + picType

		if err := c.SaveUploadedFile(file, dst); err != nil {
			_ = c.AbortWithError(http.StatusOK, apiException.ServerError)
			return
		}
		PictureList[index] = dst
	}
	utils.JsonSuccess(c, gin.H{"picture_list": PictureList})
}
