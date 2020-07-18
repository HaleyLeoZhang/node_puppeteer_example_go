package http

import (
	"github.com/gin-gonic/gin"
	"node_puppeteer_example_go/api/constant"
	"node_puppeteer_example_go/component/driver/owngin"
)

// 设置默认返回
func setDefaultResponse(c *gin.Context) {
	ownGin := &owngin.OwnGin{C: c}
	ownGin.Response(constant.HTTP_RESPONSE_CODE_SUCCESS, nil)
	c.Next()
}
