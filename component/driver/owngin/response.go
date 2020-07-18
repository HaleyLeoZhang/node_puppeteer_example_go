package owngin

import "C"
import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"node_puppeteer_example_go/api/constant"
	"time"
)

type ResponseModel struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

// HTTP 正常响应模型
func (o *OwnGin) Response(errCode int, data interface{}) {
	o.GinContext.JSON(http.StatusOK, ResponseModel{
		Code: errCode,
		Msg:  getMsg(errCode),
		Data: data,
	})
	return
}

func getMsg(code int) string {
	msg, ok := constant.HTTP_RESPONSE_MAP[code]
	if ok {
		return msg
	}

	return constant.HTTP_RESPONSE_MAP[constant.HTTP_RESPONSE_CODE_GENERAL_FAIL]
}

func NewOwnGin(c *gin.Context) *OwnGin {
	ctx, cancelFunc := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancelFunc()

	o := &OwnGin{
		GinContext: c,
		C:          ctx,
	}
	//go listenCancelContext(o)
	return o
}

func listenCancelContext(o *OwnGin) {
	select {
	case <-o.C.Done():
		errCode := getCode(o.C)
		o.Response(errCode, nil)
		o.GinContext.Abort()
	}
}

func getCode(ctx context.Context) int {
	errCode := ctx.Value(constant.HTTP_CONTEXT_GET_CODE)
	if nil == errCode {
		return constant.HTTP_RESPONSE_CODE_GENERAL_FAIL
	}
	return errCode.(int)
}
