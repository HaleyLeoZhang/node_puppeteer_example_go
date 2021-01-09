package service

import (
	"context"
	"github.com/HaleyLeoZhang/go-component/driver/xgin"
	"github.com/HaleyLeoZhang/go-component/driver/xlog"
	"github.com/HaleyLeoZhang/go-component/errgroup"
	"github.com/pkg/errors"
	"node_puppeteer_example_go/api/model"
)

func (s *Service) ChapterDetail(ctx context.Context, param *model.ChapterDetailParam) (res *model.ChapterDetailResponse, err error) {
	res = &model.ChapterDetailResponse{}
	res.Chapter = &model.ChapterDetailResponseChapter{}
	res.NextChapter = &model.ChapterDetailResponseChapter{}
	res.Comic = &model.ChapterDetailResponseComic{}
	err = nil

	chapterId := param.Id

	// 查询当前章节信息
	chapter, err := s.commonService.CurlAvatarDao.SupplierChapterGetOne(ctx, chapterId)
	if nil != err {
		return
	}
	if chapter == nil {
		err = &xgin.BusinessError{Code: xgin.HTTP_RESPONSE_CODE_SOURCE_NOT_FOUND, Message: "chapter is not exists"}
		err = errors.WithStack(err)
		return
	}
	res.Chapter.Id = chapter.Id
	res.Chapter.Name = chapter.Name
	res.Chapter.Sequence = chapter.Sequence

	eg := &errgroup.Group{}
	eg.GOMAXPROCS(2)

	supplierId := chapter.RelatedId
	supplierIdForComic := chapter.RelatedId // 防止并发读取变量
	currentSequence := chapter.Sequence

	eg.Go(func(context.Context) (errNil error) {
		errNil = nil // 一般并发业务不使用这个err返回

		nextChapter, errBusiness := s.commonService.CurlAvatarDao.SupplierChapterGetNextOne(ctx, currentSequence, supplierId)
		if nil != errBusiness {
			xlog.Errorf("ChapterDetail.Error(%+v)", err)
			return
		}
		if nextChapter == nil {
			return
		}
		res.NextChapter.Id = nextChapter.Id
		res.NextChapter.Name = nextChapter.Name
		res.NextChapter.Sequence = nextChapter.Sequence
		return
	})

	eg.Go(func(context.Context) (errNil error) {
		errNil = nil // 一般并发业务不使用这个err返回

		supplier, errBusiness := s.commonService.CurlAvatarDao.SupplierGetOne(ctx, supplierIdForComic)
		if nil != errBusiness {
			xlog.Errorf("ChapterDetail.Error(%+v)", err)
			return
		}
		if supplier == nil {
			return
		}
		comicId := supplier.RelatedId
		comic, errBusiness := s.commonService.CurlAvatarDao.ComicGetOne(ctx, comicId)
		if nil != errBusiness {
			xlog.Errorf("ChapterDetail.Error(%+v)", err)
			return
		}
		if comic == nil {
			return
		}
		res.Comic.Id = comic.Id
		res.Comic.Name = comic.Name
		return
	})
	_ = eg.Wait()
	return
}
