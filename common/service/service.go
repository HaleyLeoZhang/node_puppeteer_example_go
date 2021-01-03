package service

import (
	"github.com/HaleyLeoZhang/go-component/driver/xlog"
	"node_puppeteer_example_go/common/conf"
	"node_puppeteer_example_go/common/dao/cache"
	"node_puppeteer_example_go/common/dao/curl_avatar"
)

type Service struct {
	CurlAvatarDao *curl_avatar.Dao
	CacheDao      *cache.Dao
}

func New(cfg *conf.Config) *Service {
	s := &Service{}
	if cfg.Redis != nil {
		s.CacheDao = cache.New(cfg.Redis)
	}
	if cfg.DB != nil {
		s.CurlAvatarDao = curl_avatar.New(cfg.DB)
	}
	return s
}

// Close close the resource.
func (s *Service) Close() {
	// 各种消费者
	// - 暂无
	// 各种数据库
	// - 平滑关闭，建议数据库相关的关闭放到最后
	if s.CacheDao != nil {
		s.CacheDao.Close()
	}
	if s.CurlAvatarDao != nil {
		s.CurlAvatarDao.Close()
	}
	xlog.Info("Close.commonService.Done")
}
