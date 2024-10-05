package middleware

import (
	"StuService-Go/internal/apiException"
	"StuService-Go/pkg/utils"
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
	// 限制访问频率
	rate, err := limiter.NewRateFromFormatted("2-S")
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
