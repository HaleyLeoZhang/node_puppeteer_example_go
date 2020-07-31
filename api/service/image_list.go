package service

import (
	"context"
	"github.com/spiegel-im-spiegel/logf"
	"node_puppeteer_example_go/api/model"
)

func (s *Service) ImageList(ctx context.Context, param *model.ImageListParam) (*model.ImageListResponse, error) {
	pageId := param.PageId
	list, err := s.comicDao.GetImageList(ctx, pageId)
	if nil != err {
		logf.Errorf("ImageList.Error.%+v", err)
		return nil, err
	}
	res := &model.ImageListResponse{
		List: list,
	}
	return res, nil
}
