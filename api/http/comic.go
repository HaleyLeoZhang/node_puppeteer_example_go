package http

// ----------------------------------------------------------------------
// 漫画控制器
// ----------------------------------------------------------------------
// Link  : http://www.hlzblog.top/
// GITHUB: https://github.com/HaleyLeoZhang
// ----------------------------------------------------------------------
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"node_puppeteer_example_go/api/constant"
	"node_puppeteer_example_go/api/model"
	"node_puppeteer_example_go/component/driver/owngin"
)

type Comic struct{}

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
func (Comic) GetList(c *gin.Context) {
	ownGin := owngin.NewOwnGin(c)

	param := &model.ComicListParam{}
	err := c.Bind(param)
	if err != nil {
		ownGin.Response(constant.HTTP_RESPONSE_CODE_PARAM_INVALID, nil)
		return
	}

	res, err := srv.ComicList(ownGin.C, param)
	if nil != err {
		fmt.Printf("error: %+v", err)
	}
	ownGin.Response(constant.HTTP_RESPONSE_CODE_SUCCESS, res)
}

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
func (Comic) GetPageList(c *gin.Context) {
	ownGin := owngin.NewOwnGin(c)

	param := &model.PageListParam{}
	err := c.Bind(param)
	if err != nil {
		ownGin.Response(constant.HTTP_RESPONSE_CODE_PARAM_INVALID, nil)
		return
	}

	res, err := srv.PageList(ownGin.C, param)
	if nil != err {
		fmt.Printf("error: %+v", err)
	}
	ownGin.Response(constant.HTTP_RESPONSE_CODE_SUCCESS, res)
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
func (Comic) GetPageDetail(c *gin.Context) {
	ownGin := owngin.NewOwnGin(c)

	param := &model.PageDetailParam{}
	err := c.Bind(param)
	if err != nil {
		ownGin.Response(constant.HTTP_RESPONSE_CODE_PARAM_INVALID, nil)
		return
	}

	res, err := srv.PageDetail(ownGin.C, param)
	if nil != err {
		fmt.Printf("error: %+v", err)
	}

	ownGin.Response(constant.HTTP_RESPONSE_CODE_SUCCESS, res)
}

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
 *         "list": [
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
func (Comic) GetImageList(c *gin.Context) {
	ownGin := owngin.NewOwnGin(c)

	param := &model.ImageListParam{}
	err := c.Bind(param)
	if err != nil {
		ownGin.Response(constant.HTTP_RESPONSE_CODE_PARAM_INVALID, nil)
		return
	}

	res, err := srv.ImageList(ownGin.C, param)
	if nil != err {
		fmt.Printf("error: %+v", err)
	}

	ownGin.Response(constant.HTTP_RESPONSE_CODE_SUCCESS, res)
}
