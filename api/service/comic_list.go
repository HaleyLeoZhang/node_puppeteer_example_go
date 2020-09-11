package service

import (
	"context"
	"node_puppeteer_example_go/api/constant"
	"node_puppeteer_example_go/api/model"
	"node_puppeteer_example_go/component/driver/ownlog"
)

func (s *Service) ComicList(ctx context.Context, param *model.ComicListParam) (*model.ComicListResponse, error) {
	page := param.Page
	size := 20
	maps := make(map[string]interface{})
	maps["is_online"] = constant.TABLE_COMIC_IS_ONLINE_YES
	list, err := s.comicDao.GetComicList(ctx, page, size, maps)
	if nil != err {
		ownlog.Errorf("ComicList.Error.%+v", err)
		return nil, err
	}
	res := &model.ComicListResponse{
		List: list,
	}
	return res, nil
}
