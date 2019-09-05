package comic

// ----------------------------------------------------------------------
// 漫画章节列表-控制器
// ----------------------------------------------------------------------
// Link  : http://www.hlzblog.top/
// GITHUB: https://github.com/HaleyLeoZhang
// ----------------------------------------------------------------------

import (
	"net/http"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"github.com/HaleyLeoZhang/node_puppeteer_example_go/pkg/app"
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/pkg/e"
	// "github.com/HaleyLeoZhang/node_puppeteer_example_go/pkg/setting"
	// "github.com/HaleyLeoZhang/node_puppeteer_example_go/pkg/util"
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/service/comic_service"
	// "github.com/HaleyLeoZhang/node_puppeteer_example_go/pkg/logging"
)

/**
 * @api {get} /api/comic/page_list 漫画章节列表
 * @apiName page_list
 * @apiGroup Comic
 *
 * @apiParam {int} channel 漫画渠道ID
 * @apiParam {int} comic_id 对应渠道中的资源ID
 *
 * @apiDescription  获取漫画章节列表
 *
 * @apiVersion 1.0.0
 * @apiSuccessExample Success-Response:
 * HTTP/1.1 200 OK
 * {
 *     "code": 200,
 *     "message": "success",
 *     "data":
 *     {
 *         "list": [ // 当前数据
 *         {
 *             "id": 1,
 *             "channel": 2, // 获取漫画列表接口,对应渠道
 *             "comic_id": 5830, // 对应渠道中的资源ID
 *             "sequence": 1, // 章节序号
 *             "name": "第1话", // 章节名称
 *             "link": "https://m.manhuaniu.com/manhua/5830/200258.html", // 章节地址
 *             "progress": 2, // 章节对应图片拉取状态:枚举值：0=>未爬取，1=>处理中，2处理结束
 *             "updated_at": "2019-08-27 14:21:54",
 *             "created_at": "2019-08-27 14:22:37"
 *         }]
 *     }
 * }
 */
func GetPageList(c *gin.Context) {
	appG := app.Gin{C: c}

	channel := com.StrTo(c.Query("channel")).MustInt()
	valid := validation.Validation{}
	valid.Min(channel, 0, "channel")

	comic_id := com.StrTo(c.Query("comic_id")).MustInt()
	valid.Min(comic_id, 1, "comic_id")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	pageService := comic_service.PageParam{
		Channel: channel,
		ComicID: comic_id,
	}
	pageList, err := pageService.GetList()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.SOURCE_NOT_FOUND, nil)
		return
	}

	data := make(map[string]interface{})
	data["list"] = pageList
	// 分页逻辑放在前端,减少后端运算

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

/**
 * @api {get} /api/comic/page_detail 漫画章节详情
 * @apiName page_detail
 * @apiGroup Comic
 *
 * @apiParam {int} id 漫画章节列表接口中list对应的id
 *
 * @apiDescription 获取漫画章节详情
 *
 * @apiVersion 1.0.0
 * @apiSuccessExample Success-Response:
 * HTTP/1.1 200 OK
 */
func GetPageDetail(c *gin.Context) {
	appG := app.Gin{C: c}

	id := com.StrTo(c.Query("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	// 章节
	pageService := comic_service.PageParam{
		ID: id,
	}
	pageInfo, err := pageService.GetInfo()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.SOURCE_NOT_FOUND, nil)
		return
	}

	// 漫画
	comicService := comic_service.ComicParam{
		Channel: pageInfo.Channel,
		ComicID: pageInfo.ComicID,
	}
	comicInfo, err := comicService.GetInfo()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.SOURCE_NOT_FOUND, nil)
		return
	}

	data := make(map[string]interface{})
	data["page"] = pageInfo
	data["comic"] = comicInfo
	// 分页逻辑放在前端,减少后端运算

	appG.Response(http.StatusOK, e.SUCCESS, data)
}
