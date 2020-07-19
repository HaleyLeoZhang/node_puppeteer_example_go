package service

import (
	"context"
	"node_puppeteer_example_go/api/constant"
	"node_puppeteer_example_go/api/model"
)

func (s *Service) PageList(ctx context.Context, param *model.PageListParam) (*model.PageListResponse, error) {
	channel := param.Channel
	sourceId := param.SourceId

	maps := make(map[string]interface{})
	maps["channel"] = channel
	maps["source_id"] = sourceId
	list, err := s.comicDao.GetPageList(ctx, maps)
	if nil != err {
		ctx = context.WithValue(ctx, constant.HTTP_CONTEXT_GET_CODE, constant.HTTP_RESPONSE_CODE_GENERAL_FAIL)
		context.WithCancel(ctx)
		return nil, err
	}
	res := &model.PageListResponse{
		List: list,
	}
	return res, nil
}
