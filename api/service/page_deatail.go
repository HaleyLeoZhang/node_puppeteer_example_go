package service

import (
	"context"
	"fmt"
	"github.com/spiegel-im-spiegel/logf"
	"node_puppeteer_example_go/api/model"
	"node_puppeteer_example_go/component/errgroup"
)

func (s *Service) PageDetail(ctx context.Context, param *model.PageDetailParam) (res *model.PageDetailResponse, err error) {
	pageId := param.PageId
	currentPage, err := s.comicDao.GetPageInfo(ctx, pageId)
	if nil != err {
		logf.Errorf("PageDetail.Step1.Error.%+v", err)
		return nil, err
	}
	var nextPage *model.ComicPage
	var comic *model.Comic
	g := &errgroup.Group{}
	g.GOMAXPROCS(2)
	g.Go(func(context.Context) (err error) {
		nextPage, err = s.comicDao.GetNextPageInfo(ctx, currentPage.Channel, currentPage.SourceId, currentPage.Sequence)
		if nil != err {
			logf.Errorf("PageDetail.Step2.Error.%+v", err)
			return
		}
		return
	})
	g.Go(func(context.Context) (err error) {
		comic, err = s.comicDao.GetComicInfoWithCache(ctx, currentPage.Channel, currentPage.SourceId)
		if nil != err {
			logf.Errorf("PageDetail.Step3.Error.%+v", err)
			return
		}
		return
	})
	err = g.Wait()
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}
	res = &model.PageDetailResponse{
		Page:     currentPage,
		NextPage: nextPage,
		Comic:    comic,
	}
	return
}
