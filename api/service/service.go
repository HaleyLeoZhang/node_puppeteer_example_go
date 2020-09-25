package service

import (
	"node_puppeteer_example_go/api/conf"
	"node_puppeteer_example_go/api/dao/comic"
	//"node_puppeteer_example_go/api/dao/cache"
	"github.com/HaleyLeoZhang/go-component/driver/xlog"
)

type Service struct {
	cfg      *conf.Config
	comicDao *comic.Dao
}

// New create service instance and return.
func New(cfg *conf.Config) *Service {
	return &Service{
		cfg:      cfg,
		comicDao: comic.New(cfg),
		//cache: cache.New(cfg),  // 暂时没有业务缓存 下沉到 Dao 层
	}
}

// Close close the resource.
func (s *Service) Close() {
	s.comicDao.Close()
	xlog.Info("Close.comicDao.Done")
}
