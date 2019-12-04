package routers

import (
	// "net/http"

	"github.com/gin-gonic/gin"

	//_ "github.com/HaleyLeoZhang/node_puppeteer_example_go/docs"
	// "github.com/swaggo/gin-swagger"
	// "github.com/swaggo/gin-swagger/swaggerFiles"

	// "github.com/HaleyLeoZhang/node_puppeteer_example_go/middleware/jwt"
	// "github.com/HaleyLeoZhang/node_puppeteer_example_go/routers/api"
	// "github.com/HaleyLeoZhang/node_puppeteer_example_go/routers/api/v1"
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/routers/api/comic"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	api_comic := r.Group("/api/comic")
	api_comic.GET("comic_list", comic.GetComicList)
	api_comic.GET("page_list", comic.GetPageList)
	api_comic.GET("page_detail", comic.GetPageDetail)
	api_comic.GET("image_list", comic.GetImageList)

	// api_comic.Use(jwt.JWT())
	// {
	// 	//获取标签列表
	// 	api_comic.GET("/tags", v1.GetTags)
	// 	//新建标签
	// 	api_comic.POST("/tags", v1.AddTag)
	// 	//更新指定标签
	// 	api_comic.PUT("/tags/:id", v1.EditTag)
	// 	//删除指定标签
	// 	api_comic.DELETE("/tags/:id", v1.DeleteTag)
	// 	//导出标签
	// 	r.POST("/tags/export", v1.ExportTag)
	// 	//导入标签
	// 	r.POST("/tags/import", v1.ImportTag)

	// 	//获取文章列表
	// 	api_comic.GET("/articles", v1.GetArticles)
	// 	//获取指定文章
	// 	api_comic.GET("/articles/:id", v1.GetArticle)
	// 	//新建文章
	// 	api_comic.POST("/articles", v1.AddArticle)
	// 	//更新指定文章
	// 	api_comic.PUT("/articles/:id", v1.EditArticle)
	// 	//删除指定文章
	// 	api_comic.DELETE("/articles/:id", v1.DeleteArticle)
	// 	//生成文章海报
	// 	api_comic.POST("/articles/poster/generate", v1.GenerateArticlePoster)
	// }

	return r
}
