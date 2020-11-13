package service

import (
	"context"
	"github.com/HaleyLeoZhang/go-component/driver/xlog"
	"node_puppeteer_example_go/api/model/vo"
)

func (s *Service) PageList(ctx context.Context, param *vo.PageListParam) (*vo.PageListResponse, error) {
	channel := param.Channel
	sourceId := param.SourceId

	maps := make(map[string]interface{})
	maps["channel"] = channel
	maps["source_id"] = sourceId
	list, err := s.comicDao.GetPageList(ctx, maps)
	if nil != err {
		xlog.Errorf("PageList.Error.%+v", err)
		return nil, err
	}
	res := &vo.PageListResponse{
		List: list,
	}
	return res, nil
}
