package service

import (
	"context"
	"flag"
	"github.com/HaleyLeoZhang/go-component/driver/db"
	"github.com/HaleyLeoZhang/go-component/driver/xredis"
	"github.com/HaleyLeoZhang/go-component/errgroup"
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
		Auth:  "",
	}
	cfg.DB = &db.Config{
		Name:     "local-db",
		Type:     "mysql",
		Port:     3306,
		Database: "curl_avatar",
		User:     "",
		Password: "",
	}
	svr = New(cfg)
	ctx = context.Background()
	os.Exit(m.Run())
}

func TestService_sdasd(t *testing.T){
	//
	c := make(chan int , 2)
	eg := errgroup.Group{}
	eg.Go(func(ctx context.Context) error {
		c <- 1
	})
}