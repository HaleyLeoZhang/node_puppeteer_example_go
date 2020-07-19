package service

import (
	"context"
	"node_puppeteer_example_go/api/constant"
	"node_puppeteer_example_go/api/model"
)

func (s *Service) ComicList(ctx context.Context, param *model.ComicListParam) (*model.ComicListResponse, error) {
	page := param.Page
	size := 20
	maps := make(map[string]interface{})
	list, err := s.comicDao.GetComicList(ctx, page, size, maps)
	if nil != err {
		ctx = context.WithValue(ctx, constant.HTTP_CONTEXT_GET_CODE, constant.HTTP_RESPONSE_CODE_GENERAL_FAIL)
		context.WithCancel(ctx)
		return nil, err
	}
	res := &model.ComicListResponse{
		List: list,
	}
	return res, nil
}
