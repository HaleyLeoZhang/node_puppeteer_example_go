package service

import (
	"context"
	"node_puppeteer_example_go/api/constant"
	"node_puppeteer_example_go/api/model"
)

func (s *Service) ImageList(ctx context.Context, param *model.ImageListParam) (*model.ImageListResponse, error) {
	pageId := param.PageId
	list, err := s.comicDao.GetImageList(pageId)
	if nil != err {
		ctx = context.WithValue(ctx, constant.HTTP_CONTEXT_GET_CODE, constant.HTTP_RESPONSE_CODE_GENERAL_FAIL)
		context.WithCancel(ctx)
		return nil, err
	}
	res := &model.ImageListResponse{
		List: list,
	}
	return res, nil
}
