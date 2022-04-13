package middleware

import (
	"toolbox/logging"
	"toolbox/random"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	X_Trace_Id = "X-Trace-Id"
)

// WrapLogger wraps logger with trace id
func WrapLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get trace id from header
		traceID := c.GetHeader(X_Trace_Id)
		if traceID == "" {
			traceID = genTraceId()
		}
		zl := logging.With(zap.String("TraceId", traceID))
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
