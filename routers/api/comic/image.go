package comic

// ----------------------------------------------------------------------
// 漫画章节对应图片列表-控制器
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
 * @api {get} /api/comic/image_list 漫画章节对应图片列表
 * @apiName image_list
 * @apiGroup Comic
 *
 * @apiParam {int} page_id 漫画列表接口中list对应的id
 *
 * @apiDescription  获取漫画章节对应图片列表
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
 *                 "page_id": "1", // 漫画列表接口中list对应的id
 *                 "sequence": "1", // 图片展示序号
 *                 "src": "https://res.nbhbzl.com/images/comic/103/205259/1526284455iSc0TXw2NXnFcpd8.jpg", // 图片地址
 *                 "progress": "0", // 下载状态:枚举值:0=>待下载,1=>下载中,2下载成功
 *                 "updated_at": "2019-08-27T14:22:29+08:00",
 *                 "created_at": "2019-08-27T14:22:29+08:00"
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

	comicService := comic_service.ImageParam{
		PageID: page_id,
	}
	imageList, err := comicService.GetList()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.SOURCE_NOT_FOUND, nil)
		return
	}

	data := make(map[string]interface{})
	data["list"] = imageList
	// 分页逻辑放在前端,减少后端运算

	appG.Response(http.StatusOK, e.SUCCESS, data)
}
