package service

import (
	"context"
	"github.com/spiegel-im-spiegel/logf"
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
		logf.Errorf("PageList.Error.%+v", err)
		return nil, err
	}
	res := &model.PageListResponse{
		List: list,
	}
	return res, nil
}
