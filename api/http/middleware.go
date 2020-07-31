package http

import (
	"github.com/gin-gonic/gin"
)

// 设置默认返回
func setDefaultResponse(c *gin.Context) {
	//ownGin := &owngin.OwnGin{C: c}
	// TODO
	c.Next()
}
