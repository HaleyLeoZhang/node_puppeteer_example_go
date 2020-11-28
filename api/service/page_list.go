package service

import (
	"context"
	"github.com/HaleyLeoZhang/go-component/driver/xlog"
	"node_puppeteer_example_go/common/model/vo"
)

func (s *Service) PageList(ctx context.Context, param *vo.PageListParam) (res *vo.PageListResponse, err error) {
	channel := param.Channel
	sourceId := param.SourceId

	maps := make(map[string]interface{})
	maps["channel"] = channel
	maps["source_id"] = sourceId
	list, err := s.commonService.ComicDao.GetPageList(ctx, maps)
	if nil != err {
		xlog.Errorf("PageList.Error.%+v", err)
		return
	}
	res = &vo.PageListResponse{
		List: list,
	}
	return
}
