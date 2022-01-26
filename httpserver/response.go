package httpserver

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Response struct
type ResponseResult struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg" `
	Data  interface{} `json:"data"`
	Error interface{} `json:"error"`
}

type PageResult struct {
	Items    interface{} `json:"items"`
	Total    int         `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
}

// Response 返回的数据
func Response(c *gin.Context, code int, msg string, data interface{}, err interface{}) {
	c.JSON(code, ResponseResult{
		Code:  code,
		Msg:   msg,
		Data:  data,
		Error: err,
	})
}

// ResponseOK 返回成功
func ResponseOK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, ResponseResult{
		Code:  http.StatusOK,
		Msg:   ResponseMsg(http.StatusOK),
		Data:  data,
		Error: nil,
	})
}

// ResponseBadRequest returns 400
func ResponseBadRequest(c *gin.Context, err interface{}) {
	c.JSON(http.StatusBadRequest, ResponseResult{
		Code:  http.StatusBadRequest,
		Msg:   ResponseMsg(http.StatusBadRequest),
		Data:  nil,
		Error: err,
	})
}

// ResponseNotFound returns 400
func ResponseNotFound(c *gin.Context, err interface{}) {
	c.JSON(http.StatusBadRequest, ResponseResult{
		Code:  http.StatusNotFound,
		Msg:   ResponseMsg(http.StatusNotFound),
		Data:  nil,
		Error: err,
	})
}

// ResponseServerError returns server error
func ResponseServerError(c *gin.Context, err interface{}) {
	c.JSON(http.StatusInternalServerError, ResponseResult{
		Code:  http.StatusInternalServerError,
		Msg:   ResponseMsg(http.StatusInternalServerError),
		Data:  nil,
		Error: err,
	})
}

// ResponseCustomError returns user custom error
func ResponseCustomError(c *gin.Context, code int, err interface{}) {
	c.JSON(code, ResponseResult{
		Code:  code,
		Msg:   ResponseMsg(code),
		Data:  nil,
		Error: err,
	})
}

func HandleError(c *gin.Context, err error) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ResponseNotFound(c, err.Error())
		return
	}
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		ResponseCustomError(c, http.StatusGatewayTimeout, err.Error())
		return
	}
	ResponseServerError(c, err.Error())
}
