package middleware

import (
	"context"
	"errors"
	"net/http"
	"toolbox/httpserver"
	"toolbox/jwt"

	"github.com/gin-gonic/gin"
)

const (
	X_Access_Token = "X-Access-Token"
	ACCESS_TOKEN   = "ACCESS_TOKEN"
	UID            = "UID"
	Username       = "Username"
)

func WrapJwtTokenParser() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader(X_Access_Token)
		// get token from cookie if not found in header
		var err error
		if token == "" {
			token, err = c.Cookie(ACCESS_TOKEN)
		}
		if err != nil && !errors.Is(err, http.ErrNoCookie) {
			httpserver.ResponseUnauthorized(c, err.Error())
			c.Abort()
			return
		}
		if err != nil {
			httpserver.ResponseUnauthorized(c, "empty token")
			c.Abort()
			return
		}

		// parse token
		cliams, err := jwt.ParseToken(token, nil)
		if err != nil {
			httpserver.ResponseUnauthorized(c, err.Error())
			c.Abort()
			return
		}
		// set user id and name
		ctx := c.Request.Context()
		if cliams.UID != 0 {
			ctx = WithUID(ctx, cliams.UID)
		}
		if cliams.Username != "" {
			ctx = WithUsername(ctx, cliams.Username)
		}
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func WithUID(ctx context.Context, uid uint) context.Context {
	return context.WithValue(ctx, UID, uid)
}

func WithUsername(ctx context.Context, username string) context.Context {
	return context.WithValue(ctx, Username, username)
}

func GetUID(ctx context.Context) uint {
	if uid, ok := ctx.Value(UID).(uint); ok {
		return uid
	}
	return 0
}

func GetUsername(ctx context.Context) string {
	if username, ok := ctx.Value(Username).(string); ok {
		return username
	}
	return ""
}
