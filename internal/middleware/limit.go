package middleware

import (
	"StuService-Go/internal/apiException"
	"StuService-Go/internal/global"
	"StuService-Go/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
	"net/http"
)

type Middleware struct {
	Limiter        *limiter.Limiter
	OnError        mgin.ErrorHandler
	OnLimitReached mgin.LimitReachedHandler
	KeyGetter      mgin.KeyGetter
	ExcludedKey    func(string) bool
}

func Limit() gin.HandlerFunc {
	limitPerSec := global.Config.GetInt("Security.LimitPerSec")
	// 限制访问频率
	fmtStr := fmt.Sprintf("%d-S", limitPerSec)
	rate, err := limiter.NewRateFromFormatted(fmtStr)
	store := memory.NewStore()
	if err != nil {
		utils.Log.Fatal(err)
		return nil
	}
	middleware := mgin.NewMiddleware(limiter.New(store, rate), mgin.WithLimitReachedHandler(func(c *gin.Context) {
		_ = c.AbortWithError(http.StatusOK, apiException.LimitExceeded)
		return
	}))
	return middleware
}
