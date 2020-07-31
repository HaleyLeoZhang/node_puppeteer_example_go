package service

import (
	"context"
	"github.com/spiegel-im-spiegel/logf"
	"node_puppeteer_example_go/api/model"
	"node_puppeteer_example_go/component/driver/owngin"
)

func (s *Service) PageDetail(ctx context.Context, param *model.PageDetailParam) (*model.PageDetailResponse, error) {
	pageId := param.PageId
	currentPage, err := s.comicDao.GetPageInfo(ctx, pageId)
	if nil != err {
		logf.Errorf("PageDetail.Step1.Error.%+v", err)
		return nil, err
	}
	nextPage, err := s.comicDao.GetNextPageInfo(ctx, currentPage.Channel, currentPage.SourceId, currentPage.Sequence)
	if nil != err {
		logf.Errorf("PageDetail.Step2.Error.%+v", err)
		return nil, err
	}
	comic, err := s.comicDao.GetComicInfoWithCache(ctx, currentPage.Channel, currentPage.SourceId)
	if nil != err {
		logf.Errorf("PageDetail.Step3.Error.%+v", err)
		return nil, err
	}
	res := &model.PageDetailResponse{
		Page:     currentPage,
		NextPage: nextPage,
		Comic:    comic,
	}
	return res, nil
}
