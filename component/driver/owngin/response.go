package owngin

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"node_puppeteer_example_go/component/driver/ownlog"
	"time"
)

// ---------------------------------------------------------------------
// 		业务错误 Error 模型
// ---------------------------------------------------------------------
type BusinessError struct {
	Code    int
	Message string
	error   *error
}

func (b *BusinessError) Error() string {
	return b.Message
}

// ---------------------------------------------------------------------
// 		HTTP 响应模型
// ---------------------------------------------------------------------

type ResponseModel struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

// HTTP 响应模型
func (o *OwnGin) Response(err error, data interface{}) {
	code := HTTP_RESPONSE_CODE_SUCCESS
	message := ""
	if err != nil {
		switch err.(type) {
		case *BusinessError:
			businessError := err.(*BusinessError)
			code = businessError.Code
			message = businessError.Message
			data = nil
			ownlog.Infof("Response.BusinessError.%+v", err)
		default:
			code = HTTP_RESPONSE_CODE_UNKNOWN_FAIL
			message = "服务繁忙"
			data = nil
			ownlog.Errorf("Response.Error.%+v", err)
		}
	}
	o.GinContext.JSON(http.StatusOK, ResponseModel{
		Code: code,
		Msg:  message,
		Data: data,
	})
	return
}

func NewOwnGin(c *gin.Context) *OwnGin {
	ctx, cancelFunc := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancelFunc()

	o := &OwnGin{
		GinContext: c,
		C:          ctx,
	}
	return o
}
