package service

import (
	"context"
	"github.com/HaleyLeoZhang/go-component/driver/xlog"
	"node_puppeteer_example_go/api/constant"
	"node_puppeteer_example_go/api/model/vo"
)

func (s *Service) ComicList(ctx context.Context, param *vo.ComicListParam) (*vo.ComicListResponse, error) {
	page := param.Page
	size := 20
	maps := make(map[string]interface{})
	maps["is_online"] = constant.TABLE_COMIC_IS_ONLINE_YES
	list, err := s.comicDao.GetComicList(ctx, page, size, maps)
	if nil != err {
		xlog.Errorf("ComicList.Error.%+v", err)
		return nil, err
	}
	res := &vo.ComicListResponse{
		List: list,
	}
	return res, nil
}
