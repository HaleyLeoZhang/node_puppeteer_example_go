package service

import (
	"context"
	"flag"
	"node_puppeteer_example_go/api/conf"
	"os"
	"testing"
)

var (
	svr *Service
	ctx context.Context
)

func TestMain(m *testing.M) {
	flag.Parse()
	err := conf.Init()
	if err != nil {
		panic(err)
	}
	svr = New(conf.Conf)
	ctx = context.Background()
	ctx = context.WithValue(ctx, "page", 1)
	ctx = context.WithValue(ctx, "comic_id", 1)
	ctx = context.WithValue(ctx, "chapter_id", 1)
	os.Exit(m.Run())
}
