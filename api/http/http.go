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
		comic := &Comic{}
		/**
		 * 用户端API
		 */
		e.Group("api/comic").
			GET("comic_list", comic.GetList).
			GET("page_list", comic.GetPageList).
			GET("page_detail", comic.GetPageDetail).
			GET("image_list", comic.GetImageList)
	}
	//api_comic.GET("page_detail", comic.GetPageDetail)
	//api_comic.GET("image_list", comic.GetImageList)

	return e
}
