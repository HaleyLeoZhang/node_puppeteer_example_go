package service

import (
	"context"
	"github.com/HaleyLeoZhang/go-component/driver/xlog"
	"node_puppeteer_example_go/api/model/vo"
)

func (s *Service) ImageList(ctx context.Context, param *vo.ImageListParam) (*vo.ImageListResponse, error) {
	pageId := param.PageId
	list, err := s.comicDao.GetImageList(ctx, pageId)
	if nil != err {
		xlog.Errorf("ImageList.Error.%+v", err)
		return nil, err
	}
	res := &vo.ImageListResponse{
		List: list,
	}
	return res, nil
}
