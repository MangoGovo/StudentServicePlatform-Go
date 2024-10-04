package handler

import (
	"StuService-Go/internal/apiException"
	"StuService-Go/internal/model"
	"StuService-Go/internal/service"
	"StuService-Go/pkg/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
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

type respComment struct {
	model.Comment
	Nickname string `json:"nickname"`
}

type respFeedbackData struct {
	CreatedAt       time.Time     `json:"created_at"`
	FeedbackBy      string        `json:"feedback_by"`
	HandlerID       int64         `json:"handler_id"`
	HandlerNickname string        `json:"handler_nickname"`
	ID              int64         `json:"feedback_id"`
	FeedbackTitle   string        `json:"feedback_title"`
	FeedbackType    int           `json:"feedback_type"`
	FeedbackContent string        `json:"feedback_content"`
	FeedbackReply   []respComment `json:"feedback_reply"` // []Reply
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

	// 获取评论
	commentList, err := service.GetCommentsByFeedbackID(feedback.ID)
	if err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ServerError)
		return
	}

	feedbackReply := make([]respComment, len(commentList))

	for index, comment := range commentList {
		sender, err := service.GetUserByID(comment.SenderID)
		if err != nil {
			_ = c.AbortWithError(http.StatusOK, apiException.ServerError)
			return
		}
		feedbackReply[index] = respComment{
			Nickname: sender.Nickname,
			Comment:  comment,
		}
	}
	// 获取发帖人姓名
	feedbackBy, err := service.GetUserByID(feedback.Sender)
	if err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ServerError)
		return
	}
	var handleBy *model.User
	// 获取处理人姓名
	if feedback.Handler != 0 {
		handleBy, err = service.GetUserByID(feedback.Sender)
		if err != nil {
			_ = c.AbortWithError(http.StatusOK, apiException.ServerError)
			return
		}
	}

	// 解析图片
	var pictures []string
	picturesBytes := []byte(feedback.Pictures)
	_ = json.Unmarshal(picturesBytes, &pictures)

	respData := respFeedbackData{
		CreatedAt:       feedback.CreatedAt,
		FeedbackBy:      feedbackBy.Nickname,
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
	}

	if handleBy != nil {
		respData.HandlerID = handleBy.ID
		respData.HandlerNickname = handleBy.Nickname
	}
	utils.JsonSuccess(c, &respData)
}

type deleteFeedbackData struct {
	ID int64 `json:"feedback_id"`
}

// DeleteFeedback 删除问题反馈
func DeleteFeedback(c *gin.Context) {
	var data deleteFeedbackData
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

	if err = service.DeleteFeedback(feedback); err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ServerError)
		return
	}
	utils.JsonSuccess(c, nil)
}

type updateFeedbackData struct {
	ID              int64    `json:"feedback_id"`
	FeedbackTitle   string   `json:"feedback_title"`
	FeedbackType    int      `json:"feedback_type"`
	FeedbackContent string   `json:"feedback_content"`
	Pictures        []string `json:"pictures"`
	IsEmergency     bool     `json:"is_emergency"`
	IsAnonymous     bool     `json:"is_anonymous"`
}

// UpdateFeedback 更新问题反馈
func UpdateFeedback(c *gin.Context) {
	var data updateFeedbackData
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

	// 修改问题反馈数据
	feedback.FeedbackTitle = data.FeedbackTitle
	feedback.FeedbackType = data.FeedbackType
	feedback.FeedbackContent = data.FeedbackContent
	feedback.IsEmergency = data.IsEmergency
	feedback.IsAnonymous = data.IsAnonymous

	// 将图片列表反序列化为字符串
	pictureStr, err := utils.ConvertJsonToStr(data.Pictures)
	if err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ServerError)
		return
	}

	feedback.Pictures = pictureStr

	if err = service.UpdateFeedback(feedback); err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ServerError)
		return
	}

	utils.JsonSuccess(c, nil)
}

type getFeedbackListData struct {
	Status       int `form:"status"`
	PageCapacity int `form:"page_capacity"`
	Offset       int `form:"offset"`
}

type respFeedbackListData struct {
	CreatedAt       string `json:"created_at"`
	FeedbackID      int64  `json:"feedback_id"`
	FeedbackBy      string `json:"feedback_by"`
	FeedbackTitle   string `json:"feedback_title"`
	FeedbackType    int    `json:"feedback_type"`
	FeedbackRate    int    `json:"feedback_rate"`
	IsEmergency     bool   `json:"is_emergency"`
	FeedbackPreview string `json:"feedback_preview"`
	Status          int    `json:"status"`
}

// GetFeedbackList 获取问题反馈列表
func GetFeedbackList(c *gin.Context) {
	data := getFeedbackListData{
		Status:       -1,
		PageCapacity: 10,
		Offset:       0,
	}
	if err := c.ShouldBind(&data); err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ParamsError)
		return
	}

	user := c.MustGet("user").(*model.User)

	// 获取问题反馈列表列表
	feedbackList, err := service.GetFeedbackList(user.ID, data.Status, data.PageCapacity, data.Offset)
	if err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ServerError)
		return
	}

	// 遍历解析数据
	respFeedbackList := make([]respFeedbackListData, len(feedbackList))
	for index, feedback := range feedbackList {
		// 获取发帖者
		FeedbackBy, err := service.GetUserByID(feedback.Sender)
		if err != nil {
			_ = c.AbortWithError(http.StatusOK, apiException.ServerError)
			return
		}

		// 截取帖子内容
		content := feedback.FeedbackContent
		contentLen := len(content)
		var feedbackPreview string
		if contentLen > 50 {
			feedbackPreview = content[:50]
		} else {
			feedbackPreview = content
		}

		respFeedbackList[index] = respFeedbackListData{
			CreatedAt:       feedback.CreatedAt.Format("2006-01-02 15:04:05"),
			FeedbackID:      feedback.ID,
			FeedbackBy:      FeedbackBy.Nickname,
			FeedbackTitle:   feedback.FeedbackTitle,
			FeedbackType:    feedback.FeedbackType,
			FeedbackRate:    feedback.FeedbackRate,
			IsEmergency:     feedback.IsEmergency,
			FeedbackPreview: feedbackPreview,
			Status:          feedback.Status,
		}
	}

	total, err := service.GetFeedbackCount(user.ID, data.Status)

	if err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ServerError)
		return
	}
	utils.JsonSuccess(c, gin.H{
		"total":         total,
		"feedback_list": respFeedbackList,
	})
}

type rateFeedbackData struct {
	FeedbackID int64 `json:"feedback_id"`
	Rate       int   `json:"rate"`
}

// RateFeedback 给问题反馈打分
func RateFeedback(c *gin.Context) {
	var data rateFeedbackData
	if err := c.ShouldBind(&data); err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ParamsError)
		return
	}

	// 检查打分是否符合五分制
	rate := data.Rate
	if rate < 0 || rate > 5 {
		_ = c.AbortWithError(http.StatusOK, apiException.ParamsError)
		return
	}

	// 鉴权
	// 小细节:先判断err是否为空,为空则由于短路机制,不去处理feedback,避免了feedback为空的情况
	user := c.MustGet("user").(*model.User)
	feedback, err := service.GetFeedbackByID(data.FeedbackID)

	if err != nil || feedback.Sender != user.ID {
		_ = c.AbortWithError(http.StatusOK, apiException.PermissionsNotAllowed)
		return
	}

	//	判断问题反馈是否处于打分阶段
	if feedback.Status != 2 {
		_ = c.AbortWithError(http.StatusOK, apiException.FeedbackNotHandled)
		return
	}

	// 更新问题反馈
	feedback.FeedbackRate = data.Rate
	feedback.Status = 3
	if err = service.UpdateFeedback(feedback); err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ServerError)
		return
	}

	utils.JsonSuccess(c, nil)
}

type commentFeedbackData struct {
	FeedbackID int64  `json:"feedback_id"`
	Content    string `json:"content"`
}

// CommentFeedback 评论问题反馈
func CommentFeedback(c *gin.Context) {
	var data commentFeedbackData
	if err := c.ShouldBind(&data); err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ParamsError)
		return
	}

	// 鉴权
	// 小细节:先判断err是否为空,为空则由于短路机制,不去处理feedback,避免了feedback为空的情况
	user := c.MustGet("user").(*model.User)
	feedback, err := service.GetFeedbackByID(data.FeedbackID)

	if err != nil || feedback.Sender != user.ID {
		_ = c.AbortWithError(http.StatusOK, apiException.PermissionsNotAllowed)
		return
	}

	if err := service.CreateComment(&model.Comment{
		SenderID:   user.ID,
		FeedbackID: feedback.ID,
		Content:    data.Content,
	}); err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ServerError)
		return
	}

	utils.JsonSuccess(c, nil)
}

type deleteCommentData struct {
	FeedbackID int64 `json:"feedback_id"`
	CommentID  int64 `json:"comment_id"`
}

func DeleteComment(c *gin.Context) {
	var data deleteCommentData
	if err := c.ShouldBind(&data); err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ParamsError)
		return
	}

	// 鉴权
	// 小细节:先判断err是否为空,为空则由于短路机制,不去处理feedback,避免了feedback为空的情况
	user := c.MustGet("user").(*model.User)
	comment, err := service.GetCommentByID(data.CommentID)

	if err != nil || comment.SenderID != user.ID || comment.FeedbackID != data.FeedbackID {
		_ = c.AbortWithError(http.StatusOK, apiException.PermissionsNotAllowed)
		return
	}

	// 执行删除操作
	if err = service.DeleteComment(comment); err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ServerError)
		return
	}

	utils.JsonSuccess(c, nil)
}
