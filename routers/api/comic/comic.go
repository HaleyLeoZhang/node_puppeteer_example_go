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
 * @apiDescription  获取漫画列表[目前只用下拉模式做]
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
 *                 "id": "13",
 *                 "channel": "2",
 *                 "source_id": "12535",
 *                 "name": "我的守护女友",
 *                 "pic": "https://res.nbhbzl.com/images/cover/201909/1569035676zULh0IQhqyyh_szb.",
 *                 "intro": "末世来临，凌默的异能觉醒，他发现自己居然可以控制丧尸。 凌默利用自己的异能穿过尸潮，终于找到了心爱的...",
 *                 "weight": "1000",
 *                 "max_sequence": "0",
 *                 "updated_at": "2019-12-04 11:33:34",
 *                 "created_at": "2019-11-06 13:34:09"
 *             },
 *             {
 *                 "id": "5",
 *                 "channel": "2",
 *                 "source_id": "3181",
 *                 "name": "妖怪名单",
 *                 "pic": "https://i.loli.net/2019/09/08/sQ1Cm4vYTAViL8y.jpg",
 *                 "intro": "魅惑众生的九尾狐狸？吸人精气的合欢树妖？道家妹子求双修，仙家女神若即离。游走在这些危险分子中间可不是...",
 *                 "weight": "0",
 *                 "max_sequence": "0",
 *                 "updated_at": "2019-12-04 11:33:30",
 *                 "created_at": "2019-09-08 07:24:27"
 *             },
 *             {
 *                 "id": "7",
 *                 "channel": "2",
 *                 "source_id": "4419",
 *                 "name": "偷星九月天",
 *                 "pic": "https://res.nbhbzl.com/images/cover/201804/1524371582VN_mppKONpcP64E6.jpg",
 *                 "intro": "一场爱与梦想的奇妙冒险…… 是男仆还是热血的少年侦探？江洋大盗竟是如花美眷？！ 迷雾一层接一层，悬念...",
 *                 "weight": "0",
 *                 "max_sequence": "0",
 *                 "updated_at": "2019-12-04 11:11:39",
 *                 "created_at": "2019-09-13 12:43:58"
 *             }
 *         ]
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

	// total, err := comicService.Count()
	// if err != nil {
	// 	appG.Response(http.StatusInternalServerError, e.SOURCE_NOT_FOUND, nil)
	// 	return
	// }

	data := make(map[string]interface{})
	data["list"] = comicList
	// data["total"] = total
	// data["page"] = page
	// 分页逻辑放在前端,减少后端运算

	appG.Response(http.StatusOK, e.SUCCESS, data)
}
