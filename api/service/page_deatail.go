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
	// 填充所有默认结构体，使得前端接收结构完整
	res = &vo.PageDetailResponse{}
	res.Comic = &po.Comic{}
	res.NextPage = &po.ComicPage{}
	res.Page = &po.ComicPage{}
	err = nil

	pageId := param.PageId
	currentPage, err := s.comicDao.GetPageInfo(ctx, pageId)
	if nil != err {
		xlog.Errorf("PageDetail.Step1.Error.%+v", err)
		return nil, err
	}
	res.Page = currentPage

	g := &errgroup.Group{}
	g.GOMAXPROCS(2)
	g.Go(func(context.Context) (errNil error) {
		errNil = nil // 一般并发业务不使用这个err返回

		var errBusiness error
		res.NextPage, errBusiness = s.comicDao.GetNextPageInfo(ctx, currentPage.Channel, currentPage.SourceId, currentPage.Sequence)
		if nil != errBusiness {
			logf.Errorf("PageDetail.Step2.Error.%+v", errBusiness)
			return
		}
		return
	})
	g.Go(func(context.Context) (errNil error) {
		errNil = nil // 一般并发业务不使用这个err返回

		var errBusiness error
		res.Comic, errBusiness = s.comicDao.GetComicInfoWithCache(ctx, currentPage.Channel, currentPage.SourceId)
		if nil != errBusiness {
			logf.Errorf("PageDetail.Step3.Error.%+v", errBusiness)
			return
		}
		return
	})
	err = g.Wait()
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}
	return
}
