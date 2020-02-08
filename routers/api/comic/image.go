package comic

// ----------------------------------------------------------------------
// 漫画章节对应图片列表-控制器
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
	// "node_puppeteer_example_go/pkg/setting"
	// "node_puppeteer_example_go/pkg/util"
	"node_puppeteer_example_go/service/comic_service"
	// "node_puppeteer_example_go/pkg/logging"
)

/**
 * @api {get} /api/comic/image_list 漫画章节对应图片列表
 * @apiName image_list
 * @apiGroup Comic
 *
 * @apiParam {int} page_id 章节ID
 *
 * @apiDescription  获取漫画章节对应图片列表
 *
 * @apiSuccess {int}    code    错误码---200表示正常
 * @apiSuccess {string} message 释义---对应错误码
 * @apiSuccess {object} data    数据
 * @apiSuccess {array}  data.list  章节对应图片列表
 * @apiSuccess {int}    data.list.id  图片ID
 * @apiSuccess {int}    data.list.page_id  章节ID
 * @apiSuccess {int}    data.list.sequence 图片展示序号
 * @apiSuccess {string} data.list.src 图片地址
 * @apiSuccess {int}    data.list.progress 下载状态---枚举值:0=>待下载,1=>下载中,2下载成功
 * @apiSuccess {string} data.list.updated_at  更新时间
 * @apiSuccess {string} data.list.created_at  创建时间
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
 *                 "page_id": "1",
 *                 "sequence": "1",
 *                 "src": "https://res.nbhbzl.com/images/comic/103/205259/1526284455iSc0TXw2NXnFcpd8.jpg",
 *                 "progress": "0", //
 *                 "updated_at": "2019-08-27 14:22:29",
 *                 "created_at": "2019-08-27 14:22:29"
 *             }
 *         ]
 *     }
 * }
 */
func GetImageList(c *gin.Context) {
	appG := app.Gin{C: c}

	page_id := com.StrTo(c.Query("page_id")).MustInt()
	valid := validation.Validation{}
	valid.Min(page_id, 1, "page_id")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	imageService := comic_service.ImageParam{
		PageID: page_id,
	}
	imageList, err := imageService.GetList()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.SOURCE_NOT_FOUND, nil)
		return
	}

	data := make(map[string]interface{})
	data["list"] = imageList
	// 分页逻辑放在前端,减少后端运算

	appG.Response(http.StatusOK, e.SUCCESS, data)
}
