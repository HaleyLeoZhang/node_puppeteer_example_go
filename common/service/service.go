package service

import (
	"node_puppeteer_example_go/common/conf"
	"node_puppeteer_example_go/common/dao/comic"
	//"node_puppeteer_example_go/api/dao/cache"
	"github.com/HaleyLeoZhang/go-component/driver/xlog"
)

type Service struct {
	ComicDao *comic.Dao
}

func New(cfg *conf.Config) *Service {
	s := &Service{}
	s.ComicDao = comic.New(cfg)
	return s
}

// Close close the resource.
func (s *Service) Close() {
	// 各种消费者
	// - 暂无
	// 各种数据库
	// - 平滑关闭，建议数据库相关的关闭放到最后
	s.ComicDao.Close()
	xlog.Info("Close.commonService.comicDao.Done")
}
