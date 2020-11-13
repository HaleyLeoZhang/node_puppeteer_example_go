package service

import (
	"context"
	"github.com/HaleyLeoZhang/go-component/driver/xlog"
	"node_puppeteer_example_go/api/model/po"
	"node_puppeteer_example_go/api/model/vo"
)

func (s *Service) ImageList(ctx context.Context, param *vo.ImageListParam) (res *vo.ImageListResponse, err error) {
	pageId := param.PageId
	var list []*po.ComicImage
	list, err = s.comicDao.GetImageList(ctx, pageId)
	if nil != err {
		xlog.Errorf("ImageList.Error.%+v", err)
		return
	}
	res = &vo.ImageListResponse{
		List: list,
	}
	return
}
