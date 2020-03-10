package routers

import (
	"github.com/gin-gonic/gin"
	"node_puppeteer_example_go/routers/api/comic"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	api_comic := r.Group("/api/comic")
	api_comic.GET("comic_list", comic.GetComicList)
	api_comic.GET("page_list", comic.GetPageList)
	api_comic.GET("page_detail", comic.GetPageDetail)
	api_comic.GET("image_list", comic.GetImageList)

	return r
}
