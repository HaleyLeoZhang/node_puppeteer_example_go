package service

import (
	"context"
	"github.com/HaleyLeoZhang/go-component/driver/xlog"
	"node_puppeteer_example_go/api/constant"
	"node_puppeteer_example_go/api/model/po"
	"node_puppeteer_example_go/api/model/vo"
)

func (s *Service) ComicList(ctx context.Context, param *vo.ComicListParam) (res *vo.ComicListResponse, err error) {
	res = nil
	err = nil

	page := param.Page
	size := 20

	whereMap := make(map[string]interface{})
	whereMap["is_online"] = constant.TABLE_COMIC_IS_ONLINE_YES
	whereMap["is_deleted"] = constant.TABLE_BASE_IS_DELETED_NO

	attrMap := make(map[string]interface{})
	attrMap["offset"] = (page - 1) * size
	attrMap["order_by"] = "weight DESC,id DESC" // 权重高、新创建的在前面

	var list []*po.Comic
	list, err = s.comicDao.GetComicList(ctx, whereMap, attrMap)
	if nil != err {
		xlog.Errorf("ComicList.Error.%+v", err)
		return
	}
	res = &vo.ComicListResponse{
		List: list,
	}
	return
}
