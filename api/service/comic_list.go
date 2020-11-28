package service

import (
	"context"
	"github.com/HaleyLeoZhang/go-component/driver/xlog"
	constant2 "node_puppeteer_example_go/common/constant"
	"node_puppeteer_example_go/common/model/po"
	"node_puppeteer_example_go/common/model/vo"
)

func (s *Service) ComicList(ctx context.Context, param *vo.ComicListParam) (res *vo.ComicListResponse, err error) {
	res = nil
	err = nil

	page := param.Page
	size := 20

	whereMap := make(map[string]interface{})
	whereMap["is_online"] = constant2.TABLE_COMIC_IS_ONLINE_YES
	whereMap["is_deleted"] = constant2.TABLE_BASE_IS_DELETED_NO

	attrMap := make(map[string]interface{})
	attrMap["limit"] = size
	attrMap["offset"] = (page - 1) * size
	attrMap["order_by"] = "weight DESC,id DESC" // 权重高、新创建的在前面

	var list []*po.Comic
	list, err = s.commonService.ComicDao.GetComicList(ctx, whereMap, attrMap)
	if nil != err {
		xlog.Errorf("ComicList.Error.%+v", err)
		return
	}
	res = &vo.ComicListResponse{
		List: list,
	}
	return
}
