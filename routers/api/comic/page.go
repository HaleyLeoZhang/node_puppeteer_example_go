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
 * @apiParam {int} source_id 对应渠道中的资源ID
 *
 * @apiDescription  获取漫画章节列表
 *
 * @apiSuccess {int}    code    错误码---200表示正常
 * @apiSuccess {string} message 释义---对应错误码
 * @apiSuccess {object} data    数据
 * @apiSuccess {array}  data.list  章节列表
 * @apiSuccess {int}    data.list.id  章节ID
 * @apiSuccess {int}    data.list.channel  漫画渠道ID---获取漫画的渠道
 * @apiSuccess {int}    data.list.source_id  对应渠道中的资源ID
 * @apiSuccess {int}    data.list.sequence  章节序号
 * @apiSuccess {string} data.list.name  章节名称
 * @apiSuccess {string} data.list.link  章节地址
 * @apiSuccess {int}    data.list.progress  章节对应图片拉取状态---枚举值:0=>未爬取,1=>处理中,2处理结束
 * @apiSuccess {string} data.list.updated_at  更新时间
 * @apiSuccess {string} data.list.created_at  创建时间
 *
 * @apiVersion 1.0.0
 * @apiSuccessExample Success-Response:
 * HTTP/1.1 200 OK
 * {
 *     "code": 200,
 *     "message": "success",
 *     "data":
 *     {
 *         "list": [
 *         {
 *             "id": 1,
 *             "channel": 2,
 *             "source_id": 5830,
 *             "sequence": 1,
 *             "name": "第1话", 
 *             "link": "https://m.manhuaniu.com/manhua/5830/200258.html",
 *             "progress": 2,
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

	source_id := com.StrTo(c.Query("source_id")).MustInt()
	valid.Min(source_id, 1, "source_id")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	pageService := comic_service.PageParam{
		Channel:  channel,
		SourceID: source_id,
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
 * @apiParam {int} page_id 漫画章节列表接口中list对应的id
 *
 * @apiDescription 获取漫画章节详情
 *
 * @apiSuccess {int}    code    错误码---200表示正常
 * @apiSuccess {string} message 释义---对应错误码
 * @apiSuccess {object} data    数据
 * @apiSuccess {array}  data.comic  当前漫画信息
 * @apiSuccess {int}    data.comic.id  自增ID
 * @apiSuccess {int}    data.comic.channel  漫画渠道ID---获取漫画的渠道
 * @apiSuccess {int}    data.comic.source_id  对应渠道中的资源ID
 * @apiSuccess {string} data.comic.name  章节名称
 * @apiSuccess {string} data.comic.pic  封面图片地址
 * @apiSuccess {string} data.comic.intro  漫画简介
 * @apiSuccess {string} data.comic.updated_at  更新时间
 * @apiSuccess {string} data.comic.created_at  创建时间
 * @apiSuccess {array}  data.next_page  下一页信息
 * @apiSuccess {int}    data.next_page.id  自增ID
 * @apiSuccess {int}    data.next_page.channel  漫画渠道ID---获取漫画的渠道
 * @apiSuccess {int}    data.next_page.source_id  对应渠道中的资源ID
 * @apiSuccess {int}    data.next_page.sequence  章节序号
 * @apiSuccess {string} data.next_page.name  章节名称
 * @apiSuccess {string} data.next_page.link  章节地址
 * @apiSuccess {int}    data.next_page.progress  章节对应图片拉取状态---枚举值:0=>未爬取,1=>处理中,2处理结束
 * @apiSuccess {string} data.next_page.updated_at  更新时间
 * @apiSuccess {string} data.next_page.created_at  创建时间
 * @apiSuccess {array}  data.page  当前页信息
 * @apiSuccess {int}    data.page.id  自增ID
 * @apiSuccess {int}    data.page.channel  漫画渠道ID---获取漫画的渠道
 * @apiSuccess {int}    data.page.source_id  对应渠道中的资源ID
 * @apiSuccess {int}    data.page.sequence  章节序号
 * @apiSuccess {string} data.page.name  章节名称
 * @apiSuccess {string} data.page.link  章节地址
 * @apiSuccess {int}    data.page.progress  章节对应图片拉取状态---枚举值:0=>未爬取,1=>处理中,2处理结束
 * @apiSuccess {string} data.page.updated_at  更新时间
 * @apiSuccess {string} data.page.created_at  创建时间
 *
 * @apiVersion 1.0.0
 * @apiSuccessExample Success-Response:
 * HTTP/1.1 200 OK
 * {
 *     "code": 200,
 *     "message": "success",
 *     "data": {
 *         "comic": {
 *             "id": "1",
 *             "channel": "2",
 *             "source_id": "5830",
 *             "name": "戒魔人",
 *             "pic": "",
 *             "intro": "大一新生周小安偶然戴上一枚来历不明的商代戒指,从他口中吐出了一个恐怖的血魔人。一个人类历史上的惊天秘...",
 *             "updated_at": "2019-09-04 15:18:24",
 *             "created_at": "2019-09-03 20:37:31"
 *         },
 *         "next_page": {
 *             "id": 13,
 *             "channel": 2,
 *             "source_id": 5830,
 *             "sequence": 12,
 *             "name": "第12话",
 *             "link": "https://m.manhuaniu.com/manhua/5830/200270.html",
 *             "progress": 2,
 *             "updated_at": "2019-08-27 14:21:54",
 *             "created_at": "2019-08-27 14:24:24"
 *         },
 *         "page": {
 *             "id": 12,
 *             "channel": 2,
 *             "source_id": 5830,
 *             "sequence": 11,
 *             "name": "第11话",
 *             "link": "https://m.manhuaniu.com/manhua/5830/200269.html",
 *             "progress": 2,
 *             "updated_at": "2019-08-27 14:21:54",
 *             "created_at": "2019-08-27 14:24:17"
 *         }
 *     }
 * }
 */
func GetPageDetail(c *gin.Context) {
	appG := app.Gin{C: c}

	id := com.StrTo(c.Query("page_id")).MustInt()
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
		Channel:  pageInfo.Channel,
		SourceID: pageInfo.SourceID,
	}
	comicInfo, err := comicService.GetInfo()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.SOURCE_NOT_FOUND, nil)
		return
	}
	// 下一章节ID
	pageService.Channel = pageInfo.Channel
	pageService.SourceID = pageInfo.SourceID
	nextPageInfo, err := pageService.GetNextInfo()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.SOURCE_NOT_FOUND, nil)
		return
	}

	data := make(map[string]interface{})
	data["page"] = pageInfo
	data["next_page"] = nextPageInfo
	data["comic"] = comicInfo
	// 分页逻辑放在前端,减少后端运算

	appG.Response(http.StatusOK, e.SUCCESS, data)
}
