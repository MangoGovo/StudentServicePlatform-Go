package handler

import (
	"StuService-Go/internal/apiException"
	"StuService-Go/internal/model"
	"StuService-Go/pkg/utils"

	//"StuService-Go/internal/dao"
	"StuService-Go/internal/service"

	"github.com/gin-gonic/gin"

	//"gorm.io/gorm"
	"net/http"
)

type Amount struct {
	FeedbackAmount int64 `json:"feedback_amount"`
	UserAmount     int64 `json:"user_amount"`
	RatingAmount   []int `json:"ratings"`
}

func GetStatistics(c *gin.Context) {
	var total Amount
	feedbackSum, err := service.CountFeedback()
	if err != nil {
		c.AbortWithError(http.StatusOK, apiException.ServerError)
		return
	}
	userSum, err := service.CountUser()
	if err != nil {
		c.AbortWithError(http.StatusOK, apiException.ServerError)
		return
	}
	total.FeedbackAmount = feedbackSum
	total.UserAmount = userSum
	ratingSum, err := service.CountRating()
	if err != nil {
		/* c.AbortWithError(http.StatusOK, apiException.ServerError) */
		total.RatingAmount = []int{0, 0, 0, 0, 0}
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": total,
			"msg":  "Part NULL",
		})
		return
	}
	total.RatingAmount = ratingSum
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": total,
		"msg":  "success",
	})
}

type UserListData struct {
	UserType     int `form:"user_type"`
	PageCapacity int `form:"page_capacity"`
	Offset       int `form:"offset"`
}

type UserFinalShow struct {
	Total    int              `json:"total"`
	UserList []model.UserShow `json:"user_list"`
}

func GetUserList(c *gin.Context) {
	var data UserListData
	if err := c.ShouldBind(&data); err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ServerError)
		return
	}
	showList, err := service.GetUserList(data.UserType, data.PageCapacity, data.Offset)
	if err != nil {
		c.AbortWithError(http.StatusOK, apiException.ServerError)
		return
	}
	var userListShow UserFinalShow
	userListShow.Total = len(showList)
	userListShow.UserList = showList
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": userListShow,
	})
}

type NewUserBySudo struct {
	UserName     string `json:"username"`
	NickName     string `json:"nickname"`
	UserType     int    `json:"user_type"`
	Password     string `json:"password"`
	Gender       int    `json:"gender"`
	Introduction string `json:"introduction"`
}

func NewUser(c *gin.Context) {
	var newuser NewUserBySudo
	if err := c.ShouldBindJSON(&newuser); err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ServerError)
		return
	}
	err := service.NewUser(newuser.UserName, newuser.NickName, newuser.Password, newuser.Introduction, newuser.UserType, newuser.Gender)
	if err != nil {
		_ = c.AbortWithError(http.StatusOK, err)
		return
	}
	utils.JsonSuccess(c, nil)
}

type ChangeUserBySudo struct {
	UserID   int    `json:"user_id"`
	UserName string `json:"username"`
	NickName string `json:"nickname"`
	UserType int    `json:"user_type"`
	Password string `json:"password"`
}

func ChangeUser(c *gin.Context) {
	var changeUser ChangeUserBySudo
	if err := c.ShouldBindJSON(&changeUser); err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ServerError)
		return
	}
	if changeUser.Password != "" {
		if err := service.UpdateUser(&model.User{
			ID:       int64(changeUser.UserID),
			Username: changeUser.UserName,
			Nickname: changeUser.NickName,
			UserType: changeUser.UserType,
			Password: changeUser.Password,
		}); err != nil {
			_ = c.AbortWithError(http.StatusOK, apiException.ParamsError)
			return
		}
	} else {
		if err := service.UpdateUser(&model.User{
			ID:       int64(changeUser.UserID),
			Username: changeUser.UserName,
			Nickname: changeUser.NickName,
			UserType: changeUser.UserType,
		}); err != nil {
			_ = c.AbortWithError(http.StatusOK, apiException.ParamsError)
			return
		}
	}
	utils.JsonSuccess(c, nil)
}

type DelUserBySudo struct {
	UserID int `json:"user_id"`
}

func DelUser(c *gin.Context) {
	var deleteUser DelUserBySudo
	if err := c.ShouldBindJSON(&deleteUser); err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ServerError)
		return
	}
	if err := service.DeleteUserBySudo(deleteUser.UserID); err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ParamsError)
		return
	}
	utils.JsonSuccess(c, nil)
}

type GetRubList struct {
	PageCapacity int `form:"page_capacity"`
	Offset       int `form:"offset"`
}

type ShowRubList struct {
	Total       int              `json:"total"`
	RubbishList []model.Feedback `json:"rubbish_list"`
}

func GetRubbishList(c *gin.Context) {
	var rubList GetRubList
	if err := c.ShouldBind(&rubList); err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ServerError)
		return
	}
	rubbishList, err := service.GetRubList(rubList.PageCapacity, rubList.Offset)
	if err != nil {
		c.AbortWithError(http.StatusOK, apiException.ServerError)
		return
	}
	var rubListShow ShowRubList
	rubListShow.Total = len(rubbishList)
	rubListShow.RubbishList = rubbishList
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": rubListShow,
	})
}

type Deal struct {
	FeedbackID int  `json:"feedback_id"`
	IsRubbish  bool `json:"is_rubbish"`
}

func DealRub(c *gin.Context) {
	var dealList Deal
	if err := c.ShouldBindJSON(&dealList); err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ServerError)
		return
	}
	if err := service.DealRubbish(dealList.FeedbackID, dealList.IsRubbish); err != nil {
		_ = c.AbortWithError(http.StatusOK, apiException.ParamsError)
		return
	}
	utils.JsonSuccess(c, nil)
}
