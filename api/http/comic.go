package http

// ----------------------------------------------------------------------
// 漫画控制器
// ----------------------------------------------------------------------
// Link  : http://www.hlzbxlog.top/
// GITHUB: https://github.com/HaleyLeoZhang
// ----------------------------------------------------------------------
import (
	"fmt"
	"github.com/HaleyLeoZhang/go-component/driver/xgin"
	"github.com/HaleyLeoZhang/go-component/driver/xlog"
	"github.com/gin-gonic/gin"
	"github.com/golang/groupcache/singleflight"
	"node_puppeteer_example_go/api/model"
)

type Comic struct{}

/**
 * @api {get} /api/comic/list 漫画列表
 * @apiName list
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
 * @apiSuccess {string} data.list.name  漫画名称
 * @apiSuccess {string} data.list.pic  封面图片地址
 * @apiSuccess {string} data.list.intro  漫画简介
 * @apiSuccess {int} data.list.weight  权重值.值越大,越靠前展示
 * @apiSuccess {int} data.list.tag  标记。枚举值: 0:没有标记,1:热门,2:连载,3:完结
 * @apiSuccess {string} data.list.supplier  当前绑定的渠道信息
 * @apiSuccess {int} data.list.supplier.id  渠道ID
 * @apiSuccess {int} data.list.supplier.max_sequence  对应渠道当前的最大阅读序号
 *
 * @apiVersion 1.0.0
 * @apiSuccessExample Success-Response:
 * {
 *     "code": 200,
 *     "message": "",
 *     "data": {
 *         "list": [
 *             {
 *                 "id": 1,
 *                 "name": "百炼成神",
 *                 "pic": "https://res1.xiaoqinre.com/images/cover/201807/1530935442Y6Tc2lgwA6XJPVum.jpg",
 *                 "intro": "漫画简介：现在身为卑微家奴的罗征，本身家中大少爷，因家族败落，妹妹被强大势力囚禁，无奈只得听命于人。可是天无绝人之路，父亲留给他的古书中竟然暗藏炼器神法，可将人炼制成器！而隐藏在这背后的神秘力量到底是什么？这是一场与命运的较量。",
 *                 "weight": 200,
 *                 "tag": 2,
 *                 "supplier": {
 *                     "id": 1,
 *                     "max_sequence": 663
 *                 }
 *             }
 *         ]
 *     }
 * }
 */
func (Comic) GetList(c *gin.Context) {
	xGin := xgin.NewGin(c)

	param := &model.ComicListParam{}
	err := c.Bind(param)
	if err != nil {
		err = &xgin.BusinessError{Code: xgin.HTTP_RESPONSE_CODE_PARAM_INVALID, Message: "Param is invalid"}
		xGin.Response(err, nil)
		return
	}

	// 幂等请求，防止击穿 说明文档 https://segmentfault.com/a/1190000018464029
	g := &singleflight.Group{}
	groupKey := fmt.Sprintf("comic_list_%v", param.Page)
	res, err := g.Do(groupKey, func() (data interface{}, errBusiness error) {
		data, errBusiness = srv.ComicList(xGin.C, param)
		return
	})
	xGin.Response(err, res)
}

/**
 * @api {get} /api/chapter/list 漫画章节列表
 * @apiName list
 * @apiGroup Chapter
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
 * @apiSuccess {int}    data.list.sequence  章节顺序号
 * @apiSuccess {string}    data.list.name  章节名
 *
 * @apiVersion 1.0.0
 * @apiSuccessExample Success-Response:
 * {
 *     "code": 200,
 *     "message": "",
 *     "data": {
 *         "list": [
 *             {
 *                 "id": 1,
 *                 "sequence": 1,
 *                 "name": "第1话 炼器功法"
 *             }
 *         ]
 *     }
 * }
 */
func (Comic) GetChapterList(c *gin.Context) {
	xGin := xgin.NewGin(c)

	param := &model.ChapterListParam{}
	err := c.Bind(param)
	if err != nil {
		err = &xgin.BusinessError{Code: xgin.HTTP_RESPONSE_CODE_PARAM_INVALID, Message: "Param is invalid"}
		xGin.Response(err, nil)
		return
	}
	// 幂等请求，防止击穿 说明文档 https://segmentfault.com/a/1190000018464029
	g := &singleflight.Group{}
	groupKey := fmt.Sprintf("chapter_list_%v", param.ComicId)
	res, err := g.Do(groupKey, func() (data interface{}, errBusiness error) {
		data, errBusiness = srv.ChapterList(xGin.C, param)
		return
	})
	xGin.Response(err, res)
}

/**
 * @api {get} /api/chapter/detail 漫画章节详情
 * @apiName detail
 * @apiGroup Chapter
 *
 * @apiParam {int} page_id 漫画章节列表接口中list对应的id
 *
 * @apiDescription 获取漫画章节详情
 *
 * @apiSuccess {int}    code    错误码---200表示正常
 * @apiSuccess {string} message 释义---对应错误码
 * @apiSuccess {object} data    数据
 * @apiSuccess {array}  data.comic  当前漫画信息
 * @apiSuccess {int}    data.comic.id  漫画ID
 * @apiSuccess {string} data.comic.name  漫画名
 * @apiSuccess {string} data.comic.intro  漫画简介
 * @apiSuccess {array}  data.next_chapter  下一章节信息
 * @apiSuccess {int}    data.next_chapter.id  章节ID
 * @apiSuccess {string} data.next_chapter.name  章节名
 * @apiSuccess {int}    data.next_chapter.sequence  章节序号
 * @apiSuccess {array}  data.chapter  当前章节信息
 * @apiSuccess {int}    data.chapter.id  章节ID
 * @apiSuccess {string} data.chapter.name  章节名
 * @apiSuccess {int}    data.chapter.sequence  章节序号
 *
 * @apiVersion 1.0.0
 * @apiSuccessExample Success-Response:
 * {
 *     "code": 200,
 *     "message": "",
 *     "data": {
 *         "chapter": {
 *             "id": 1,
 *             "sequence": 1,
 *             "name": "第1话 炼器功法"
 *         },
 *         "next_chapter": {
 *             "id": 2,
 *             "sequence": 2,
 *             "name": "第2话 族炼日"
 *         },
 *         "comic": {
 *             "id": 1,
 *             "name": "百炼成神",
 *             "intro": "漫画简介：现在身为卑微家奴的罗征，本身家中大少爷，因家族败落，妹妹被强大势力囚禁，无奈只得听命于人。可是天无绝人之路，父亲留给他的古书中竟然暗藏炼器神法，可将人炼制成器！而隐藏在这背后的神秘力量到底是什么？这是一场与命运的较量。"
 *         }
 *     }
 * }
 */
func (Comic) GetChapterDetail(c *gin.Context) {
	xGin := xgin.NewGin(c)

	param := &model.ChapterDetailParam{}
	err := c.Bind(param)
	if err != nil {
		err = &xgin.BusinessError{Code: xgin.HTTP_RESPONSE_CODE_PARAM_INVALID, Message: "Param is invalid"}
		xGin.Response(err, nil)
		return
	}

	// 幂等请求，防止击穿 说明文档 https://segmentfault.com/a/1190000018464029
	g := &singleflight.Group{}
	groupKey := fmt.Sprintf("chapter_detail_%v", param.Id)
	res, err := g.Do(groupKey, func() (data interface{}, errBusiness error) {
		data, errBusiness = srv.ChapterDetail(xGin.C, param)
		return
	})

	xGin.Response(err, res)
}

/**
 * @api {get} /api/curl_avatar/image_list 漫画章节对应图片列表
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
	xGin := xgin.NewGin(c)

	param := &model.ImageListParam{}
	err := c.Bind(param)
	if err != nil {
		err = &xgin.BusinessError{Code: xgin.HTTP_RESPONSE_CODE_PARAM_INVALID, Message: "Param is invalid"}
		xGin.Response(err, nil)
		return
	}

	var res *model.ImageListResponse
	// 幂等请求，防止击穿 说明文档 https://segmentfault.com/a/1190000018464029
	g := &singleflight.Group{}
	groupKey := fmt.Sprintf("image_list_%v", param.PageId)
	g.Do(groupKey, func() (interface{}, error) {
		res, err = srv.ImageList(xGin.C, param)
		if nil != err {
			xlog.Errorf("router.GetList.Err.%+v", err)
		}
		return res, err
	})
	xGin.Response(err, res)
}
