package service

import (
	"context"
	"github.com/HaleyLeoZhang/go-component/driver/xgin"
	"github.com/HaleyLeoZhang/go-component/driver/xlog"
	"github.com/HaleyLeoZhang/go-component/errgroup"
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/api/model"
)

func (s *Service) ChapterDetail(ctx context.Context, param *model.ChapterDetailParam) (res *model.ChapterDetailResponse, err error) {
	res = &model.ChapterDetailResponse{
		Chapter:     &model.ChapterDetailResponseChapter{},
		NextChapter: &model.ChapterDetailResponseChapter{},
		Comic:       &model.ChapterDetailResponseComic{},
	}
	chapterId := param.Id
	// 2022年9月14日 17:24:53 预期当前页面缓存命中率低于50%，暂不做接口级缓存

	// 查询当前章节信息
	chapter, err := s.commonService.CurlAvatarDao.SupplierChapterGetOne(ctx, chapterId)
	if nil != err {
		return
	}
	if chapter == nil {
		err = &xgin.BusinessError{Code: xgin.HTTP_RESPONSE_CODE_SOURCE_NOT_FOUND, Message: "chapter is not exists"}
		return
	}
	res.Chapter = &model.ChapterDetailResponseChapter{
		Id:       chapter.Id,
		Name:     chapter.Name,
		Sequence: chapter.Sequence,
	}

	eg := &errgroup.Group{}
	eg.GOMAXPROCS(2)

	var (
		supplierId         = chapter.RelatedId
		supplierIdForComic = chapter.RelatedId // 防止并发读取变量
		currentSequence    = chapter.Sequence
	)
	eg.Go(func(context.Context) (errNil error) {
		nextChapter, errBusiness := s.commonService.CurlAvatarDao.SupplierChapterGetNextOne(ctx, currentSequence, supplierId)
		if nil != errBusiness {
			xlog.Errorf("ChapterDetail.Error(%+v)", errBusiness)
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
		supplier, errBusiness := s.commonService.CurlAvatarDao.SupplierGetOne(ctx, supplierIdForComic)
		if nil != errBusiness {
			xlog.Errorf("ChapterDetail.Error(%+v)", errBusiness)
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
		res.Comic = &model.ChapterDetailResponseComic{
			Id:   comic.Id,
			Name: comic.Name,
		}
		return
	})
	err = eg.Wait()
	if err != nil {
		return
	}
	return
}
