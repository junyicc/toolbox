package middleware

import (
	"time"
	"toolbox/logging"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// WrapTimer times the process of the request
func WrapTimer() gin.HandlerFunc {
	return func(c *gin.Context) {
		l := logging.GetContextLogger(c.Request.Context())
		start := time.Now()
		l.Info("request in",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("remote addr", c.Request.RemoteAddr),
			zap.String("token", c.Request.Header.Get("X-Access-Token")),
		)
		c.Next()
		l.Info("response out",
			zap.Int("status", c.Writer.Status()),
			zap.Int("size", c.Writer.Size()),
			zap.Int64("duration(ms)", time.Since(start).Milliseconds()),
		)
	}
}
