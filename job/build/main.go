package main

import (
	"flag"
	"github.com/HaleyLeoZhang/go-component/driver/bootstrap"
	"github.com/HaleyLeoZhang/go-component/driver/xlog"
	"node_puppeteer_example_go/api/conf"
	"node_puppeteer_example_go/api/service"
)

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(err)
	}

	oneService := service.New(conf.Conf)
	xlog.Init(conf.Conf.Log)

	app := bootstrap.New()
	app.Start(func() { // 此部分代码，请勿阻塞进程
		// 通知错误
		//
		return
	}).Stop(func() {
		oneService.Close()
	})

}
