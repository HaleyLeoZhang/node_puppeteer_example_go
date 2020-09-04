package service

import (
	"context"
	"node_puppeteer_example_go/api/model"
	"node_puppeteer_example_go/component/driver/ownlog"
)

func (s *Service) ImageList(ctx context.Context, param *model.ImageListParam) (*model.ImageListResponse, error) {
	pageId := param.PageId
	list, err := s.comicDao.GetImageList(ctx, pageId)
	if nil != err {
		ownlog.Errorf("ImageList.Error.%+v", err)
		return nil, err
	}
	res := &model.ImageListResponse{
		List: list,
	}
	return res, nil
}
