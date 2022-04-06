package httpserver

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	ErrNotFound      = errors.New("not found")
	ErrAlreadyExists = errors.New("already exists")
	ErrInvalidData   = errors.New("invalid data")
	ErrUnauthorized  = errors.New("unauthorized")
	ErrForbidden     = errors.New("forbidden")
)

func HandleError(c *gin.Context, err error) {
	if errors.Is(err, ErrNotFound) || errors.Is(err, gorm.ErrRecordNotFound) {
		ResponseNotFound(c, err.Error())
		return
	}
	if errors.Is(err, ErrAlreadyExists) {
		ResponseBadRequest(c, err.Error())
		return
	}
	if errors.Is(err, ErrInvalidData) || errors.Is(err, gorm.ErrInvalidData) || errors.Is(err, gorm.ErrInvalidField) {
		ResponseBadRequest(c, err.Error())
		return
	}
	if errors.Is(err, ErrUnauthorized) {
		ResponseUnauthorized(c, err.Error())
		return
	}
	if errors.Is(err, ErrForbidden) {
		ResponseForbidden(c, err.Error())
		return
	}
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		ResponseCustomError(c, http.StatusGatewayTimeout, err.Error())
		return
	}
	ResponseServerError(c, err.Error())
}
