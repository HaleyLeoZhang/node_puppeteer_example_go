package service

import (
	"github.com/HaleyLeoZhang/go-component/driver/xconsul"
	"github.com/HaleyLeoZhang/go-component/driver/xlog"
	"github.com/pkg/errors"
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/common/conf"
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/common/dao/cache"
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/common/dao/curl_avatar"
)

type Service struct {
	CurlAvatarDao *curl_avatar.Dao
	CacheDao      *cache.Dao

	microService *xconsul.Client
}

func New(cfg *conf.Config) *Service {
	s := &Service{}
	if cfg.Redis != nil {
		s.CacheDao = cache.New(cfg.Redis)
	}
	if cfg.DB != nil {
		s.CurlAvatarDao = curl_avatar.New(cfg.DB)
	}
	// 服务注册
	if cfg.Consul.Addr != "" {
		var err error
		s.microService, err = xconsul.NewClient(&cfg.Consul)
		if err != nil {
			panic(errors.WithStack(err))
		}
		err = s.microService.HttpRegister()
		if err != nil {
			panic(errors.WithStack(err))
		}
	}
	return s
}

// Close close the resource.
func (s *Service) Close() {
	// 注销服务
	if s.microService != nil {
		_ = s.microService.Deregister()
	}
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
