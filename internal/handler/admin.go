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

type orderData struct {
	FeedbackID int64 `json:"feedback_id"`
}

func Order(c *gin.Context) {
	var data orderData
	if err := c.ShouldBind(&data); err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ParamsError)
		return
	}

	// 判断是否可以被接单
	feedback, err := service.GetFeedbackByID(data.FeedbackID)

	if err != nil || feedback.Status != 0 {
		_ = c.AbortWithError(http.StatusOK, apiException.PermissionsNotAllowed)
		return
	}
	user := c.MustGet("user").(*model.User)

	feedback.Status = 1
	feedback.Handler = user.ID

	if err = service.UpdateFeedback(feedback); err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ServerError)
		return
	}

	utils.JsonSuccess(c, nil)
}

type undoOrderData struct {
	FeedbackID int64 `json:"feedback_id"`
}

func UndoOrder(c *gin.Context) {
	var data undoOrderData
	if err := c.ShouldBind(&data); err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ParamsError)
		return
	}

	// 判断是否可以被撤单
	user := c.MustGet("user").(*model.User)
	feedback, err := service.GetFeedbackByID(data.FeedbackID)

	if err != nil || (feedback.Status != 1 && feedback.Status != 2) || feedback.Handler != user.ID {
		_ = c.AbortWithError(http.StatusOK, apiException.PermissionsNotAllowed)
		return
	}

	feedback.Status = 0
	feedback.Handler = 0

	if err = service.UpdateFeedback(feedback); err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ServerError)
		return
	}

	utils.JsonSuccess(c, nil)
}

type rubbishData struct {
	FeedbackID int64 `json:"feedback_id"`
}

func Rubbish(c *gin.Context) {
	var data rubbishData
	if err := c.ShouldBind(&data); err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ParamsError)
		return
	}

	feedback, err := service.GetFeedbackByID(data.FeedbackID)
	if err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.PermissionsNotAllowed)
		return
	}
	feedback.IsRubbish = 1
	if err = service.UpdateFeedback(feedback); err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ServerError)
		return
	}

	utils.JsonSuccess(c, nil)
}

// AdminComment 管理员评论问题反馈
func AdminComment(c *gin.Context) {
	var data commentFeedbackData
	if err := c.ShouldBind(&data); err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ParamsError)
		return
	}

	// 鉴权
	user := c.MustGet("user").(*model.User)
	feedback, err := service.GetFeedbackByID(data.FeedbackID)

	if err != nil || (feedback.Sender != user.ID && feedback.Handler != user.ID) {
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
	feedback.Status = 2
	if err := service.UpdateFeedback(feedback); err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ServerError)
		return
	}
	utils.JsonSuccess(c, nil)
}

type adminDeleteCommentData struct {
	FeedbackID int64 `json:"feedback_id"`
	CommentID  int64 `json:"comment_id"`
}

func AdminDelComment(c *gin.Context) {
	var data adminDeleteCommentData
	if err := c.ShouldBind(&data); err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ParamsError)
		return
	}

	// 鉴权
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
	// 删除后判断是否还有留存的回复
	feedback, err := service.GetFeedbackByID(data.FeedbackID)
	if err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ServerError)
		return
	}

	total, err := service.GetCommentCount(feedback.ID)
	if err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ServerError)
		return
	}

	if total == 0 {
		// 恢复到接单但未回复
		feedback.Status = 1
		if err = service.UpdateFeedback(feedback); err != nil {
			_ = c.AbortWithError(http.StatusOK, apiException.ServerError)
			return
		}
	}
	utils.JsonSuccess(c, nil)

}

func AdminGetFeedbackList(c *gin.Context) {
	data := getFeedbackListData{
		Status:       -1,
		PageCapacity: 10,
		Offset:       0,
	}
	if err := c.ShouldBind(&data); err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ParamsError)
		return
	}

	//user := c.MustGet("user").(*model.User)

	// 获取问题反馈列表列表
	feedbackList, err := service.AdminGetFeedbackList(data.Status, data.PageCapacity, data.Offset)
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
			continue
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

	total, err := service.AdminGetFeedbackCount(data.Status)

	if err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ServerError)
		return
	}
	utils.JsonSuccess(c, gin.H{
		"total":         total,
		"feedback_list": respFeedbackList,
	})
}

func AdminQueryFeedback(c *gin.Context) {
	var data queryFeedbackData
	if err := c.ShouldBind(&data); err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ParamsError)
		return
	}

	// 鉴权
	// 小细节:先判断err是否为空,为空则由于短路机制,不去处理feedback,避免了feedback为空的情况
	//user := c.MustGet("user").(*model.User)
	feedback, err := service.GetFeedbackByID(data.ID)
	if err != nil {
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
			continue
		}
		feedbackReply[index] = respComment{
			Nickname: sender.Nickname,
			UserType: sender.UserType,
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
		handleBy, err = service.GetUserByID(feedback.Handler)
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
		SenderID:        feedbackBy.ID,
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
