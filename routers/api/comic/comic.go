package comic

// ----------------------------------------------------------------------
// 漫画列表-控制器
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
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/pkg/setting"
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/pkg/util"
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/service/comic_service"
	// "github.com/HaleyLeoZhang/node_puppeteer_example_go/pkg/logging"
)

/**
 * @api {get} /api/comic/comic_list 漫画列表
 * @apiName comic_list
 * @apiGroup Comic
 *
 * @apiParam {int} page 页码
 *
 * @apiDescription  获取漫画列表
 *
 * @apiVersion 1.0.0
 * @apiSuccessExample Success-Response:
 * HTTP/1.1 200 OK
 * {
 *     "code": 200,
 *     "message": "success",
 *     "data": {
 *         "list": [ // 当前数据
 *             {
 *                 "id": "1",
 *                 "channel": "2", // 枚举值：0.未知 1.腾讯漫画 2.漫画牛
 *                 "comic_id": "5830", // 对应渠道中的资源ID
 *                 "name": "戒魔人", // 漫画名称
 *                 "pic": "", // 漫画封面
 *                 "intro": "大一新生周小安偶然戴上一枚来历不明的商代戒指，从他口中吐出了一个恐怖的血魔人。一个人类历史上的惊天秘...", // 漫画简介
 *                 "updated_at": "2019-08-27T14:20:02+08:00",
 *                 "created_at": "2019-09-03T20:37:31+08:00"
 *             }
 *         ],
 *         "page": 1, // 当前页码
 *         "total": 3 // 数据条数
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
