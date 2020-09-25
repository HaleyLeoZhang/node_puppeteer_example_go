package service

import (
	"context"
	"github.com/HaleyLeoZhang/go-component/driver/xlog"
	"node_puppeteer_example_go/api/model"
)

func (s *Service) ImageList(ctx context.Context, param *model.ImageListParam) (*model.ImageListResponse, error) {
	pageId := param.PageId
	list, err := s.comicDao.GetImageList(ctx, pageId)
	if nil != err {
		xlog.Errorf("ImageList.Error.%+v", err)
		return nil, err
	}
	res := &model.ImageListResponse{
		List: list,
	}
	return res, nil
}
