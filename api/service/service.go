package service

import (
	"github.com/HaleyLeoZhang/go-component/driver/xlog"
	"node_puppeteer_example_go/api/conf"
	comonnconf "node_puppeteer_example_go/common/conf"
	commonservice "node_puppeteer_example_go/common/service"
)

type Service struct {
	commonService *commonservice.Service
}

// New create service instance and return.
func New(cfg *conf.Config) *Service {
	s := &Service{}
	cfgCommon := &comonnconf.Config{}
	cfgCommon.DB = cfg.DB
	cfgCommon.Redis = cfg.Redis
	s.commonService = commonservice.New(cfgCommon)
	return s
}

// Close close the resource.
func (s *Service) Close() {
	// 各种消费者
	// - 暂无
	// 各种数据库
	// - 平滑关闭，建议数据库相关的关闭放到最后
	s.commonService.Close()
	xlog.Info("Close.Service.Done")
}
