package service

import (
	"context"
	"flag"
	"github.com/HaleyLeoZhang/go-component/driver/db"
	"github.com/HaleyLeoZhang/go-component/driver/xredis"
	"node_puppeteer_example_go/common/conf"
	"os"
	"testing"
)

var (
	svr *Service
	ctx context.Context
)

func TestMain(m *testing.M) {
	flag.Parse()
	cfg := &conf.Config{}
	cfg.Redis = &xredis.Config{
		Name:  "local—redis",
		Proto: "tcp",
		Addr:  "192.168.56.110:6379",
		Auth:  "zhangli",
	}
	cfg.DB = &db.Config{
		Name:     "local-db",
		Type:     "mysql",
		Port:     3306,
		Database: "curl_avatar",
		User:     "yth_blog",
		Password: "http://hlzblog.top",
	}
	svr = New(cfg)
	ctx = context.Background()
	os.Exit(m.Run())
}
