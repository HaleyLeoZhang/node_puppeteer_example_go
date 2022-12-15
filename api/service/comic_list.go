package service

import (
	"context"
	"encoding/json"
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/api/model"
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/common/model/po"
)

func (s *Service) ComicList(ctx context.Context, param *model.ComicListParam) (res *model.ComicListResponse, err error) {
	res, err = s.commonService.CacheDao.ComicListRequestWithCache(ctx, param.GetPage(), func() (by []byte, errRaw error) {
		// 生成缓存前查询
		resRaw, errRaw := s.comicList(ctx, param)
		if errRaw != nil {
			return
		}
		by, _ = json.Marshal(resRaw)
		return
	})
	return
}

func (s *Service) comicList(ctx context.Context, param *model.ComicListParam) (res *model.ComicListResponse, err error) {
	res = &model.ComicListResponse{
		List: make([]*model.ComicListResponseItem, 0),
	}
	var (
		page = param.GetPage()
		size = param.GetPageSize()
		// -
		limit       = size
		offset      = (page - 1) * size
		orderBy     = "weight DESC,id DESC" // 权重高、新创建的在前面
		comicFields = "id,name,intro,pic,tag,weight,related_id"
	)
	// 查询漫画列表
	comicList, err := s.commonService.CurlAvatarDao.ComicListForIndexWithFields(ctx, limit, offset, orderBy, comicFields)
	if nil != err {
		return
	}
	// 查询渠道数据
	if len(comicList) == 0 {
		return
	}
	var (
		supplierIds    = make([]int, 0, len(comicList))
		supplierFields = "id,max_sequence"
	)
	for _, comic := range comicList {
		supplierIds = append(supplierIds, comic.RelatedId)
	}
	supplierList, err := s.commonService.CurlAvatarDao.SupplierListForIndexWithFields(ctx, supplierIds, supplierFields)
	if nil != err {
		return
	}
	lenSupplierList := len(supplierList)
	if lenSupplierList == 0 {
		return
	}
	// 渠道列表转map
	mapSupplier := make(map[int]*po.Supplier)
	for _, supplier := range supplierList {
		mapSupplier[supplier.Id] = supplier
	}
	// 漫画查询渠道
	res.List = make([]*model.ComicListResponseItem, 0, lenSupplierList) // 没有关联渠道的就不展示了
	for _, comic := range comicList {
		if supplier, ok := mapSupplier[comic.RelatedId]; ok {
			tmp := &model.ComicListResponseItem{}
			tmp.Id = comic.Id
			tmp.Name = comic.Name
			tmp.Intro = comic.Intro
			tmp.Pic = comic.Pic
			tmp.Tag = comic.Tag
			tmp.Weight = comic.Weight
			tmpSupplier := &model.ComicListResponseSupplier{}
			tmpSupplier.Id = supplier.Id
			tmpSupplier.MaxSequence = supplier.MaxSequence
			tmp.Supplier = tmpSupplier
			res.List = append(res.List, tmp)
		}
	}
	return
}
