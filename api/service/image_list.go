package service

import (
	"context"
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/api/model"
)

func (s *Service) ImageList(ctx context.Context, param *model.ImageListParam) (res *model.ImageListResponse, err error) {
	res = &model.ImageListResponse{
		List: make([]*model.ImageListResponseItem, 0),
	}
	// 先找渠道
	var (
		chapterId = param.ChapterId
		// -
		supplierFields = "src_origin,src_own,sequence"
	)
	supplierImageList, err := s.commonService.CurlAvatarDao.SupplierImageListWithFields(ctx, chapterId, supplierFields)
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
