package service

import (
	"context"
	"github.com/HaleyLeoZhang/go-component/driver/xgin"
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/api/model"
)

func (s *Service) ChapterList(ctx context.Context, param *model.ChapterListParam) (res *model.ChapterListResponse, err error) {
	res = &model.ChapterListResponse{
		List: make([]*model.ChapterListResponseItem, 0),
	}
	// 先找渠道
	var (
		comicId = param.ComicId
		// -
		supplierFields = "id"
	)
	supplier, err := s.commonService.CurlAvatarDao.SupplierOneForChapterWithFields(ctx, comicId, supplierFields)
	if err != nil {
		return
	}
	if supplier == nil {
		err = xgin.NewBusinessError("暂无可用渠道", xgin.HTTP_RESPONSE_CODE_SOURCE_NOT_FOUND)
		return
	}
	var (
		supplierChapterFields = "id,name,sequence"
	)
	// 再找渠道对应章节信息
	supplierChapterList, err := s.commonService.CurlAvatarDao.SupplierChapterListWithFields(ctx, supplier.Id, supplierChapterFields)
	if nil != err {
		return
	}
	lenSupplierChapterList := len(supplierChapterList)
	if lenSupplierChapterList == 0 {
		return
	}
	for _, supplierChapter := range supplierChapterList {
		res.List = append(res.List, &model.ChapterListResponseItem{
			Id:       supplierChapter.Id,
			Name:     supplierChapter.Name,
			Sequence: supplierChapter.Sequence,
		})
	}

	return
}
