package service

import (
	"context"
	"github.com/HaleyLeoZhang/go-component/driver/xlog"
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/api/model"
	constant2 "github.com/HaleyLeoZhang/node_puppeteer_example_go/common/constant"
)

func (s *Service) ChapterList(ctx context.Context, param *model.ChapterListParam) (res *model.ChapterListResponse, err error) {
	comicId := param.ComicId

	res = &model.ChapterListResponse{}
	res.List = make([]*model.ChapterListResponseItem, 0)
	err = nil

	// 先找渠道
	whereSupplierMap := make(map[string]interface{})
	whereSupplierMap["related_id"] = comicId
	whereSupplierMap["status"] = constant2.BASE_TABLE_ONLINE

	attrSupplierMap := make(map[string]interface{})
	attrSupplierMap["limit"] = 1
	attrSupplierMap["offset"] = 0
	attrSupplierMap["order_by"] = "weight DESC,id ASC" // 权重高、先创建的在前面

	supplierList, err := s.commonService.CurlAvatarDao.SupplierList(ctx, whereSupplierMap, attrSupplierMap)
	if nil != err {
		xlog.Errorf("PageList.Error.%+v", err)
		return
	}
	lenSupplierList := len(supplierList)
	if lenSupplierList == 0 {
		return
	}
	supplier := supplierList[0]
	// 再找渠道对应章节信息
	whereSupplierChapterMap := make(map[string]interface{})
	whereSupplierChapterMap["related_id"] = supplier.Id
	whereSupplierChapterMap["status"] = constant2.BASE_TABLE_ONLINE

	attrSupplierChapterMap := make(map[string]interface{})
	attrSupplierChapterMap["order_by"] = "sequence ASC" // 权重高、先创建的在前面

	supplierChapterList, err := s.commonService.CurlAvatarDao.SupplierChapterList(ctx, whereSupplierChapterMap, attrSupplierChapterMap)
	if nil != err {
		return
	}
	lenSupplierChapterList := len(supplierChapterList)
	if lenSupplierChapterList == 0 {
		return
	}
	for _, supplierChapter := range supplierChapterList {
		tmp := &model.ChapterListResponseItem{}
		tmp.Id = supplierChapter.Id
		tmp.Name = supplierChapter.Name
		tmp.Sequence = supplierChapter.Sequence
		res.List = append(res.List, tmp)
	}

	return
}
