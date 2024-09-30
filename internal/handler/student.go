package handler

import (
	"StuService-Go/internal/apiException"
	"StuService-Go/internal/model"
	"StuService-Go/internal/service"
	"StuService-Go/pkg/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type postFeedbackData struct {
	FeedbackTitle   string   `json:"feedback_title"`
	FeedbackType    int      `json:"feedback_type"`
	FeedbackContent string   `json:"feedback_content"`
	Pictures        []string `json:"pictures"`
	IsEmergency     bool     `json:"is_emergency"`
	IsAnonymous     bool     `json:"is_anonymous"`
}

// PostFeedback 发布问题反馈
func PostFeedback(c *gin.Context) {
	var data postFeedbackData
	if err := c.ShouldBindJSON(&data); err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ParamsError)
		return
	}
	// 将json转化为字符串,便于存储在MySQL中
	PicturesStr, err := utils.ConvertJsonToStr(data.Pictures)

	if err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ParamsError)
		return
	}

	user := c.MustGet("user").(*model.User)

	if err = service.CreateFeedback(&model.Feedback{
		FeedbackTitle:   data.FeedbackTitle,
		FeedbackType:    data.FeedbackType,
		FeedbackContent: data.FeedbackContent,
		Pictures:        PicturesStr,
		IsEmergency:     data.IsEmergency,
		IsAnonymous:     data.IsAnonymous,
		Sender:          user.ID,
	}); err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ServerError)
		return
	}
	utils.JsonSuccess(c, nil)
}

type queryFeedbackData struct {
	ID int64 `form:"feedback_id"`
}
type respFeedbackData struct {
	ID              int64         `json:"feedback_id"`
	FeedbackTitle   string        `json:"feedback_title"`
	FeedbackType    int           `json:"feedback_type"`
	FeedbackContent string        `json:"feedback_content"`
	FeedbackReply   []model.Reply `json:"feedback_reply"` // []Reply
	Status          int           `json:"status"`
	FeedbackRate    int           `json:"feedback_rate"`
	Pictures        []string      `json:"pictures"` // []string
	IsEmergency     bool          `json:"is_emergency"`
	IsAnonymous     bool          `json:"is_anonymous"`
}

// QueryFeedback 查询问题反馈
func QueryFeedback(c *gin.Context) {
	var data queryFeedbackData
	if err := c.ShouldBind(&data); err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ParamsError)
		return
	}

	// 鉴权
	// 小细节:先判断err是否为空,为空则由于短路机制,不去处理feedback,避免了feedback为空的情况
	user := c.MustGet("user").(*model.User)
	feedback, err := service.GetFeedbackByID(data.ID)
	if err != nil || feedback.Sender != user.ID {
		_ = c.AbortWithError(http.StatusOK, apiException.PermissionsNotAllowed)
		return
	}

	// 解析评论
	var feedbackReply []model.Reply
	feedbackReplyBytes := []byte(feedback.FeedbackReply)

	// 解析不出来就为空+
	if err = json.Unmarshal(feedbackReplyBytes, &feedbackReply); err != nil {
		feedbackReply = []model.Reply{}
	}

	// 解析图片
	var pictures []string
	picturesBytes := []byte(feedback.Pictures)
	_ = json.Unmarshal(picturesBytes, &pictures)
	utils.JsonSuccess(c, &respFeedbackData{
		ID:              feedback.ID,
		FeedbackTitle:   feedback.FeedbackTitle,
		FeedbackType:    feedback.FeedbackType,
		FeedbackContent: feedback.FeedbackContent,
		FeedbackReply:   feedbackReply,
		Status:          feedback.Status,
		FeedbackRate:    feedback.FeedbackRate,
		Pictures:        pictures,
		IsEmergency:     feedback.IsEmergency,
		IsAnonymous:     feedback.IsAnonymous,
	})
}

// DeleteFeedback 删除问题反馈
func DeleteFeedback(c *gin.Context) {

}

// UpdateFeedback 更新问题反馈
func UpdateFeedback(c *gin.Context) {

}

// GetFeedbackList 获取问题反馈列表
func GetFeedbackList(c *gin.Context) {

}
