package http

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/groupcache/singleflight"
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/api/conf"
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/api/service"
)

var srv *service.Service
var g = &singleflight.Group{} // 接口级缓存 幂等请求，防止击穿 说明文档 https://segmentfault.com/a/1190000018464029

func Init(e *gin.Engine, srvInjection *service.Service) *gin.Engine {
	srv = srvInjection

	//e.Use() // 暂无中间件需要被设置
	{
		comic := &Comic{}
		/**
		 * 用户端API
		 * TODO 接口级缓存
		 */
		routComic := e.Group("api/")
		routComic.GET("comic/list", comic.GetList)
		routComic.GET("chapter/list", comic.GetChapterList)
		routComic.GET("chapter/detail", comic.GetChapterDetail)
		routComic.GET("image/list", comic.GetImageList)
	}
	// 健康检查专用
	{
		e.GET(conf.Conf.Consul.HealthCheckRouter, func(c *gin.Context) {
			// 健康检测啥都不用处理
			return
		})

	}

	return e
}
