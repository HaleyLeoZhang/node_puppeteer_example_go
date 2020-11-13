package service

import (
	"context"
	"fmt"
	"github.com/HaleyLeoZhang/go-component/driver/xlog"
	"github.com/HaleyLeoZhang/go-component/errgroup"
	"github.com/spiegel-im-spiegel/logf"
	"node_puppeteer_example_go/api/model/po"
	"node_puppeteer_example_go/api/model/vo"
)

func (s *Service) PageDetail(ctx context.Context, param *vo.PageDetailParam) (res *vo.PageDetailResponse, err error) {
	pageId := param.PageId
	currentPage, err := s.comicDao.GetPageInfo(ctx, pageId)
	if nil != err {
		xlog.Errorf("PageDetail.Step1.Error.%+v", err)
		return nil, err
	}
	var nextPage *po.ComicPage
	var comic *po.Comic
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
	res = &vo.PageDetailResponse{
		Page:     currentPage,
		NextPage: nextPage,
		Comic:    comic,
	}
	return
}
