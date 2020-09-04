package owngin

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

// 超时时间
var timeout time.Duration

// 在原来的包上，增加功能
type OwnGin struct {
	GinContext *gin.Context // owngin context
	C          context.Context
}

type Config struct {
	Name    string        `yaml:"name" json:"name"` // 用于 Trace 识别
	Debug   bool          `yaml:"debug" json:"debug"`
	Timeout time.Duration `yaml:"timeout" json:"timeout"`
}

func New(c *Config) *gin.Engine {
	gin.SetMode("release")
	if true == c.Debug {
		gin.SetMode("debug")
	}
	timeout = c.Timeout

	r := gin.New()
	r.Use(gin.Logger())
	//r.Use(gin.Recovery()) // 外部 recovery
	return r
}
