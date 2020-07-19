package service

import (
	"context"
	"node_puppeteer_example_go/api/constant"
	"node_puppeteer_example_go/api/model"
)

func (s *Service) PageDetail(ctx context.Context, param *model.PageDetailParam) (*model.PageDetailResponse, error) {
	pageId := param.PageId
	currentPage, err := s.comicDao.GetPageInfo(ctx, pageId)
	if nil != err {
		ctx = context.WithValue(ctx, constant.HTTP_CONTEXT_GET_CODE, constant.HTTP_RESPONSE_CODE_GENERAL_FAIL)
		context.WithCancel(ctx)
		return nil, err
	}
	nextPage, err := s.comicDao.GetNextPageInfo(ctx, currentPage.Channel, currentPage.SourceId, currentPage.Sequence)
	if nil != err {
		ctx = context.WithValue(ctx, constant.HTTP_CONTEXT_GET_CODE, constant.HTTP_RESPONSE_CODE_GENERAL_FAIL)
		context.WithCancel(ctx)
		return nil, err
	}
	comic, err := s.comicDao.GetComicInfoWithCache(ctx, currentPage.Channel, currentPage.SourceId)
	if nil != err {
		ctx = context.WithValue(ctx, constant.HTTP_CONTEXT_GET_CODE, constant.HTTP_RESPONSE_CODE_GENERAL_FAIL)
		context.WithCancel(ctx)
		return nil, err
	}
	res := &model.PageDetailResponse{
		Page:     currentPage,
		NextPage: nextPage,
		Comic:    comic,
	}
	return res, nil
}
