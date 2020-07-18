package http

import (
	"github.com/gin-gonic/gin"
	"node_puppeteer_example_go/api/conf"
	"node_puppeteer_example_go/api/service"
)

var srv *service.Service

func Init(c *conf.Config, e *gin.Engine) *gin.Engine {
	srv = service.New(c)

	//e.Use(setDefaultResponse)

	{
		/**
		 * 用户端API
		 */
		// 老接口重构
		// - 这一组接口下的术语定义 收藏、订阅 文档 http://wiki.shihuo.cn/pages/viewpage.action?pageId=20261307
		e.Group("api/comic").
			GET("comic_list", Comic{}.GetList).
			GET("page_list", Comic{}.GetPageList).
			GET("page_detail", Comic{}.GetPageDetail).
			GET("image_list", Comic{}.GetImageList)
	}
	//api_comic.GET("page_detail", comic.GetPageDetail)
	//api_comic.GET("image_list", comic.GetImageList)

	return e
}
