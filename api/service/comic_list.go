package service

import (
	"context"
	"encoding/json"
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/api/model"
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/common/constant"
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

	page := param.GetPage()
	size := param.GetPageSize()

	whereComicMap := make(map[string]interface{})
	whereComicMap["status"] = constant.BASE_TABLE_ONLINE

	attrComicMap := make(map[string]interface{})
	attrComicMap["limit"] = size
	attrComicMap["offset"] = (page - 1) * size
	attrComicMap["order_by"] = "weight DESC,id DESC" // 权重高、新创建的在前面

	// 查询漫画列表
	comicList, err := s.commonService.CurlAvatarDao.ComicList(ctx, whereComicMap, attrComicMap)
	if nil != err {
		return
	}
	// 查询渠道数据
	lenComicList := len(comicList)
	if lenComicList == 0 {
		return
	}
	comicIdList := make([]int, 0, lenComicList)
	for _, comic := range comicList {
		comicIdList = append(comicIdList, comic.RelatedId)
	}
	whereSupplierMap := make(map[string]interface{})
	whereSupplierMap["status"] = constant.BASE_TABLE_ONLINE
	supplierList, err := s.commonService.CurlAvatarDao.SupplierList(ctx, whereSupplierMap, nil)
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
