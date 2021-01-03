package http

import (
	"github.com/gin-gonic/gin"
	"node_puppeteer_example_go/api/service"
)

var srv *service.Service

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

	return e
}
