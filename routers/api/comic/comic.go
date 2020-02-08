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
 * @apiSuccess {string} data.list.max_sequence  对应漫画当前的最大阅读序号。对应comic_pages.sequence的对应最大值
 * @apiSuccess {string} data.list.weight  权重值.值越大,越靠前展示
 * @apiSuccess {string} data.list.tag  标记。枚举值: 0->没有标记,1->热门,2->连载,3->完结
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
 *                 "id": "9",
 *                 "channel": "3",
 *                 "source_id": "11406",
 *                 "name": "天才高手",
 *                 "pic": "https://www.onemanhua.com/comic/11406/cover.jpg",
 *                 "intro": "不定期更新傲娇女总裁，娇羞小秘书，南非黑钻公主，特种兵女保镖，和爱吃胡萝卜的屌丝美男会擦出怎样的火花是宿命纠缠还是心机暗算，走进天才高手，带你领略不一样的助理人生。 ",
 *                 "max_sequence": "186",
 *                 "weight": "100",
 *                 "tag": "0",
 *                 "updated_at": "2019-12-04 15:58:48",
 *                 "created_at": "2019-12-04 11:52:37"
 *             },
 *             {
 *                 "id": "1",
 *                 "channel": "2",
 *                 "source_id": "5830",
 *                 "name": "戒魔人",
 *                 "pic": "https://i.loli.net/2019/09/08/czSNHV3fnyaox65.jpg",
 *                 "intro": "大一新生周小安偶然戴上一枚来历不明的商代戒指，从他口中吐出了一个恐怖的血魔人。一个人类历史上的惊天秘...",
 *                 "max_sequence": "0",
 *                 "weight": "0",
 *                 "tag": "0",
 *                 "updated_at": "2019-09-14 06:20:16",
 *                 "created_at": "2019-09-03 12:37:31"
 *             },
 *             {
 *                 "id": "2",
 *                 "channel": "2",
 *                 "source_id": "10660",
 *                 "name": "一人之下",
 *                 "pic": "https://i.loli.net/2019/09/05/F4nyW9iHltuK6Ur.jpg",
 *                 "intro": "随着爷爷尸体被盗，神秘少女冯宝宝的造访，少年张楚岚的平静校园生活被彻底颠覆。急于解开爷爷和自身秘密的...",
 *                 "max_sequence": "0",
 *                 "weight": "0",
 *                 "tag": "0",
 *                 "updated_at": "2019-09-14 06:20:16",
 *                 "created_at": "2019-09-03 12:37:33"
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
