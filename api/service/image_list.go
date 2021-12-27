package service

import (
	"context"
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/api/model"
	constant2 "github.com/HaleyLeoZhang/node_puppeteer_example_go/common/constant"
)

func (s *Service) ImageList(ctx context.Context, param *model.ImageListParam) (res *model.ImageListResponse, err error) {
	chapterId := param.ChapterId
	res = &model.ImageListResponse{}
	res.List = make([]*model.ImageListResponseItem, 0)

	whereSupplierImageMap := make(map[string]interface{})
	whereSupplierImageMap["related_id"] = chapterId
	whereSupplierImageMap["status"] = constant2.BASE_TABLE_ONLINE

	attrSupplierImageMap := make(map[string]interface{})
	attrSupplierImageMap["order_by"] = "sequence ASC" // 权重高、先创建的在前面

	supplierImageList, err := s.commonService.CurlAvatarDao.SupplierImageList(ctx, whereSupplierImageMap, attrSupplierImageMap)
	if nil != err {
		return
	}
	lenSupplierImageList := len(supplierImageList)
	if lenSupplierImageList == 0 {
		return
	}
	for _, supplierImage := range supplierImageList {
		tmp := &model.ImageListResponseItem{}
		tmp.SrcOrigin = supplierImage.SrcOrigin
		tmp.SrcOwn = supplierImage.SrcOwn
		tmp.Sequence = supplierImage.Sequence
		res.List = append(res.List, tmp)
	}

	return
}
