package service

import (
	"context"
	"flag"
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/api/conf"
	"os"
	"testing"
)

var (
	svr *Service
	ctx context.Context
)

func TestMain(m *testing.M) {
	_ = flag.Set("conf", "../build/app.yaml")
	err := conf.Init()
	if err != nil {
		panic(err)
	}
	svr = New(conf.Conf)
	ctx = context.Background()
	os.Exit(m.Run())
}
