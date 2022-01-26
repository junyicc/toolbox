package middleware

import (
	"toolbox/logging"
	"toolbox/random"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// WrapLogger wraps logger with trace id
func WrapLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		zl := logging.With(zap.String("TraceId", genTraceId()))
		// AddCallerSkip(-1) 保证调用者的信息能正确输出
		zl.SetLogger(zl.GetLogger().WithOptions(zap.AddCallerSkip(-1)))

		reqCtx := logging.WithValue(c.Request.Context(), zl)
		c.Request = c.Request.WithContext(reqCtx)
		c.Next()
	}
}

func genTraceId() string {
	return random.RandString(21)
}
