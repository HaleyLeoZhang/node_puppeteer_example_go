package service

import (
	"context"
	"node_puppeteer_example_go/api/model"
)

func (s *Service) ImageList(ctx context.Context, param *model.ImageListParam) (res *model.ImageListResponse, err error) {
	//pageId := param.PageId
	//var list []*po.ComicImage
	//list, err = s.commonService.ComicDao.GetImageList(ctx, pageId)
	//if nil != err {
	//	xlog.Errorf("ImageList.Error.%+v", err)
	//	return
	//}
	//res = &model.ImageListResponse{
	//	List: list,
	//}
	return
}
