package comic

// ----------------------------------------------------------------------
// 漫画列表-控制器
// ----------------------------------------------------------------------
// Link  : http://www.hlzblog.top/
// GITHUB: https://github.com/HaleyLeoZhang
// ----------------------------------------------------------------------

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"node_puppeteer_example_go/pkg/app"
	"node_puppeteer_example_go/pkg/e"
	"node_puppeteer_example_go/pkg/setting"
	"node_puppeteer_example_go/pkg/util"
	"node_puppeteer_example_go/service/comic_service"
	// "node_puppeteer_example_go/pkg/logging"
)

/**
 * @api {get} /api/comic/comic_list 漫画列表
 * @apiName comic_list
 * @apiGroup Comic
 *
 * @apiParam {int} page 页码
 *
 * @apiDescription  获取漫画列表[目前只是用下拉模式做的,后续可能会引入普通分页]
 *
 * @apiSuccess {int}    code    错误码---200表示正常
 * @apiSuccess {string} message 释义---对应错误码
 * @apiSuccess {object} data    数据
 * @apiSuccess {array}  data.list  漫画列表
 * @apiSuccess {int}    data.list.id  漫画ID
 * @apiSuccess {int}    data.list.channel  漫画渠道ID---枚举值:0.未知 1.腾讯漫画 2.漫画牛
 * @apiSuccess {int}    data.list.source_id  对应渠道中的资源ID
 * @apiSuccess {string} data.list.name  章节名称
 * @apiSuccess {string} data.list.pic  封面图片地址
 * @apiSuccess {string} data.list.intro  漫画简介
 * @apiSuccess {string} data.list.updated_at  更新时间
 * @apiSuccess {string} data.list.created_at  创建时间
 * @apiSuccess {int}    data.page  当前页码
 * @apiSuccess {int}    data.total  数据条数
 *
 * @apiVersion 1.0.0
 * @apiSuccessExample Success-Response:
 * HTTP/1.1 200 OK
 * {
 *     "code": 200,
 *     "message": "success",
 *     "data": {
 *         "list": [
 *             {
 *                 "id": "1",
 *                 "channel": "2",
 *                 "source_id": "5830",
 *                 "name": "戒魔人",
 *                 "pic": "",
 *                 "intro": "大一新生周小安偶然戴上一枚来历不明的商代戒指,从他口中吐出了一个恐怖的血魔人。一个人类历史上的惊天秘...",
 *                 "updated_at": "2019-08-27 14:20:02",
 *                 "created_at": "2019-09-03 20:37:31"
 *             }
 *         ],
 *         "page": 1,
 *         "total": 3
 *     }
 * }
 */
func GetComicList(c *gin.Context) {
	appG := app.Gin{C: c}

	page := com.StrTo(c.Query("page")).MustInt()
	// logging.Info("page %s", page)
	valid := validation.Validation{}
	valid.Min(page, 1, "page")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	comicService := comic_service.ComicParam{
		PageNum:  util.GetPage(c),
		PageSize: setting.AppSetting.PageSize,
	}
	comicList, err := comicService.GetList()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.SOURCE_NOT_FOUND, nil)
		return
	}

	total, err := comicService.Count()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.SOURCE_NOT_FOUND, nil)
		return
	}

	data := make(map[string]interface{})
	data["list"] = comicList
	data["total"] = total
	data["page"] = page
	// 分页逻辑放在前端,减少后端运算

	appG.Response(http.StatusOK, e.SUCCESS, data)
}
