package service

import (
	"context"
	"github.com/spiegel-im-spiegel/logf"
	"node_puppeteer_example_go/api/model"
)

func (s *Service) ComicList(ctx context.Context, param *model.ComicListParam) (*model.ComicListResponse, error) {
	page := param.Page
	size := 20
	maps := make(map[string]interface{})
	list, err := s.comicDao.GetComicList(ctx, page, size, maps)
	if nil != err {
		logf.Errorf("ComicList.Error.%+v", err)
		return nil, err
	}
	res := &model.ComicListResponse{
		List: list,
	}
	return res, nil
}
